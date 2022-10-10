package cli

import (
	"github.com/urfave/cli/v2"
)

var traceFlag = cli.BoolFlag{
	Name:  "trace",
	Usage: "enable trace mode",
}
var debugFlag = cli.BoolFlag{
	Name:    "debug",
	Aliases: []string{"d"},
	Usage:   "enable debug mode",
}

func App() *cli.App {
	return &cli.App{
		Name:  "omnia",
		Usage: "some tools for nexxOmnia",
		Action: func(ctx *cli.Context) error {
			cli.ShowAppHelp(ctx)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:   "import",
				Usage:  "import data from pyton script",
				Action: importCmd,
				Flags: []cli.Flag{
					&debugFlag,
				},
			},
			{
				Name:   "merge",
				Usage:  "merge data into omnia",
				Action: mergeCmd,
				Flags: []cli.Flag{
					&debugFlag,
				},
			},
		},
	}
}
