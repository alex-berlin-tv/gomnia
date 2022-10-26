package cli

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia"
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/enums"
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/params"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func completeCmd(ctx *cli.Context) error {
	if ctx.Bool("trace") {
		log.SetLevel(log.TraceLevel)
	} else if ctx.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}
	data := collectionFromFile(ctx.String("input"))
	client := omnia.OmniaFromFile(ctx.String("config"))

	rsl, err := client.EditableAttributes(enums.AudioStreamType)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%+v", rsl.Result)

	return nil
	for i, item := range data {
		log.Infof("[%d/%d] Update Metadata of %s", i+1, len(data), item.FileName)
		client.Update(enums.AudioStreamType, item.Id, params.Custom{
			"title":       item.Title,
			"description": item.Description,
		})
	}

	return nil
}
