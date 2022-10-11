package cli

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const omniaFile = "omnia.json"
const dataFile = "data/old.json"

func mergeCmd(ctx *cli.Context) error {
	if ctx.Bool("trace") {
		log.SetLevel(log.TraceLevel)
	} else if ctx.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}
	client := omnia.OmniaFromFile(omniaFile)
	client.Call("get", omnia.AudioType, "all", []string{}, &omnia.QueryParameters{
		AddPublishingDetails: omnia.YesBool,
	})
	return nil
}
