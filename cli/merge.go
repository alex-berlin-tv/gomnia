package cli

import (
	"strings"

	"github.com/alex-berlin-tv/nexx_omnia_go/omnia"
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/enums"
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/params"

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
	data := collectionFromFile(ctx.String("input"))
	client := omnia.OmniaFromFile(omniaFile)

	for i, entry := range data {
		log.Infof("[%d/%d] Query ID for %s", i+1, len(data), entry.FileName)
		data[i].Id = getIdForTitle(client, entry.FileName)
	}

	data.toFile(ctx.String("output"))

	// _, err := client.All(omnia.AudioStreamType, &omnia.BasicParameters{
	// 	AddPublishingDetails: omnia.YesBool,
	// 	Limit:                100,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Info(rsl)
	// client.ById(omnia.AudioStreamType, 967567, nil)
	// client.Call("get", omnia.AudioStreamType, "all", []string{}, &omnia.BasicParameters{
	// AddPublishingDetails: omnia.YesBool,
	// })
	// client.ManagementCall("put", omnia.AudioStreamType, "update", []string{"967567"}, omnia.CustomParameters{
	// 	"title": "fnord",
	// })
	return nil
}

func getIdForTitle(client omnia.Omnia, title string) int {
	title = strings.Split(title, ".")[0]
	rsl, err := client.ByQuery(enums.AudioStreamType, title, &params.ByQueryParameters{
		BasicParameters: params.BasicParameters{
			AddPublishingDetails: enums.YesBool,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if len(rsl.Result) != 1 {
		log.Error("There are %d result for the %s query", len(rsl.Result), title)
		return -1
	}
	return rsl.Result[0].General.Id
}
