package cli

import (
	"context"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

func copyCoverCmd(ctx *cli.Context) error {
	collection := collectionFromFile(ctx.String("input"))

	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	var lastCoverFile string
	for data := range ch {
		if string(data) == lastCoverFile {
			continue
		}

		id, err := strconv.Atoi(string(data))
		if err != nil {
			log.Warnf("No cover for id %s found", string(data))
			continue
		}
		entry, err := collection.byId(id)
		if err != nil {
			log.Warnf("No cover for id %s found", string(data))
			continue
		}
		clipboard.Write(clipboard.FmtText, []byte(entry.Image))
		log.Infof("cover for id %d is %s", id, entry.Image)
		lastCoverFile = entry.Image
	}
	return nil
}
