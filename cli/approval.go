package cli

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia"
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/enums"
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/params"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// All stuff regarding the approval and publishing process.

func approveCmd(ctx *cli.Context) error {
	client := omnia.OmniaFromFile(ctx.String("config"))
	ids, err := getIds(ctx)
	if err != nil {
		log.Fatal(err)
	}
	streamType, err := getEnum[enums.StreamType](ctx, "type")
	if err != nil {
		log.Panic(err)
	}
	rsl, err := client.Approve(*streamType, ids[0], params.Approve{})
	if err != nil {
		log.Error(err)
	}
	log.Info(rsl)
	return nil
}

func publishCmd(ctx *cli.Context) error {
	client := omnia.OmniaFromFile(ctx.String("config"))
	log.Info(client)
	return nil
}

func rejectCmd(ctx *cli.Context) error {
	client := omnia.OmniaFromFile(ctx.String("config"))
	log.Info(client)
	return nil
}

func unblockCmd(ctx *cli.Context) error {
	client := omnia.OmniaFromFile(ctx.String("config"))
	log.Info(client)
	return nil
}

func unpublishCmd(ctx *cli.Context) error {
	client := omnia.OmniaFromFile(ctx.String("config"))
	log.Info(client)
	return nil
}
