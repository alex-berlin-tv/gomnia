package main

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/cli"

	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cli.App().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
