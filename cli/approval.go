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
		log.Fatal(err)
	}
	for _, id := range ids {
		rsl, err := client.Approve(*streamType, id, params.Approve{})
		if err != nil {
			handleApiError(rsl, err, false)
		}
		if ctx.Bool("publish") {
			rsl, err := client.Publish(*streamType, id)
			if err != nil {
				handleApiError(rsl, err, false)
			}
		}
	}
	return nil
}

func publishCmd(ctx *cli.Context) error {
	client := omnia.OmniaFromFile(ctx.String("config"))
	ids, err := getIds(ctx)
	if err != nil {
		log.Fatal(err)
	}
	streamType, err := getEnum[enums.StreamType](ctx, "type")
	if err != nil {
		log.Fatal(err)
	}
	for _, id := range ids {
		rsl, err := client.Publish(*streamType, id)
		if err != nil {
			handleApiError(rsl, err, false)
		}
	}
	return nil
}

func rejectCmd(ctx *cli.Context) error {
	client := omnia.OmniaFromFile(ctx.String("config"))
	ids, err := getIds(ctx)
	if err != nil {
		log.Fatal(err)
	}
	streamType, err := getEnum[enums.StreamType](ctx, "type")
	if err != nil {
		log.Fatal(err)
	}
	for _, id := range ids {
		rsl, err := client.Reject(*streamType, id, params.Reject{})
		if err != nil {
			handleApiError(rsl, err, false)
		}
	}
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
