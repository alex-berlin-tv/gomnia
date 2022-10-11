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
	client.ById(omnia.AudioStreamType, 967567, nil)
	// client.Call("get", omnia.AudioStreamType, "all", []string{}, &omnia.BasicParameters{
	// AddPublishingDetails: omnia.YesBool,
	// })
	// client.ManagementCall("put", omnia.AudioStreamType, "update", []string{"967567"}, omnia.CustomParameters{
	// 	"title": "fnord",
	// })
	return nil
}
