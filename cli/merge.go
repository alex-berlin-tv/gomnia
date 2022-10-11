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
	rsl, err := client.ByQuery(omnia.AudioStreamType, "Detektei Zukunft #5: Chatbots und Community Management Made in Brandenburg", &omnia.BasicParameters{
		AddPublishingDetails: omnia.YesBool,
	})
	if err != nil {
		log.Error(err)
	}
	// client.ById(omnia.AudioStreamType, 967567, nil)
	// client.Call("get", omnia.AudioStreamType, "all", []string{}, &omnia.BasicParameters{
	// AddPublishingDetails: omnia.YesBool,
	// })
	// client.ManagementCall("put", omnia.AudioStreamType, "update", []string{"967567"}, omnia.CustomParameters{
	// 	"title": "fnord",
	// })
	log.Info(rsl.Result)
	return nil
}
