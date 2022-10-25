package cli

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var filesToImport = []string{
	"data/detektei.csv",
	"data/machtwas.csv",
	"data/millionen.csv",
	"data/spotlight.csv",
	"data/transit.csv",
}
var outPath = "data/old.json"

type entry struct {
	FileName    string `json:"file_name"`
	Title       string
	Date        time.Time
	Description string
	Image       string
	Url         string
	Id          int
}

func entryFromReader(rec []string) entry {
	date, err := time.Parse("02.01.2006", rec[2])
	if err != nil {
		log.Fatal(err)
	}
	return entry{
		FileName:    rec[0],
		Title:       rec[1],
		Date:        date,
		Description: rec[3],
		Image:       rec[4],
		Url:         rec[5],
	}
}

type collection []entry

func collectionFromFile(path string) collection {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var rsl collection
	if json.Unmarshal([]byte(file), &rsl); err != nil {
		log.Fatal(err)
	}
	return rsl
}

func (c collection) toFile(path string) {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile(path, data, 0644); err != nil {
		log.Fatal(err)
	}
}

func (c collection) byId(id int) (*entry, error) {
	for _, entry := range c {
		if entry.Id == id {
			return &entry, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("no entry found for id %d", id))
}

func importCmd(ctx *cli.Context) error {
	var rsl collection
	for _, path := range filesToImport {
		file, _ := os.Open(path)
		reader := csv.NewReader(file)
		firstDone := false
		for {
			rec, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			if firstDone {
				rsl = append(rsl, entryFromReader(rec))
			} else {
				firstDone = true
			}
		}
	}
	file, _ := json.MarshalIndent(rsl, "", "  ")
	_ = ioutil.WriteFile(outPath, file, 0644)
	return nil
}
