package cli

import (
	"github.com/urfave/cli/v2"
)

func App() *cli.App {
	return &cli.App {
		Name: "omnia",
		Usage: "some tools for nexxOmnia",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "debug",
				Aliases: []string{"d"},
				Usage: "enable debug mode",
			},
		},
		Action: func(ctx *cli.Context) error {
			cli.ShowAppHelp(ctx)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "merge",
				Usage: "merge data into omnia",
				Action: mergeCmd,
			},
		},
	}
}
