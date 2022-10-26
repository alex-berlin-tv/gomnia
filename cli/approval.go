package cli

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// All stuff regarding the approval and publishing process.

func approveCmd(ctx *cli.Context) error {
	ids, err := getIds(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(ids)
	// client := omnia.OmniaFromFile(ctx.String("config"))
	// log.Info(client)
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
