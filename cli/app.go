package cli

import (
	"fmt"
	"strings"

	"github.com/alex-berlin-tv/nexx_omnia_go/omnia"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// Sets up the logging.
func setupLogging(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}
	if ctx.Bool("trace") {
		log.SetLevel(log.TraceLevel)
	}
	return nil
}

type enumFlag[T ~string] struct {
	Default  omnia.Enum[T]
	selected T
}

func (e *enumFlag[T]) Set(value string) error {
	rsl, err := omnia.EnumByValue(e.Default, T(value))
	if err != nil {
		return err
	}
	e.selected = *rsl
	return nil
}

func (e enumFlag[T]) String() string {
	return string(e.selected)
}

func (e enumFlag[T]) fmtPossibleValues() string {
	return fmt.Sprintf("[%s]", strings.Join(omnia.EnumValues(e.Default), ", "))
}

func (e enumFlag[T]) fmtUsage(desc string) string {
	return fmt.Sprintf("%s %s", desc, e.fmtPossibleValues())
}

var configFileFlag = cli.StringFlag{
	Name:    "config",
	Aliases: []string{"c"},
	Usage:   "path to config file",
}

var streamTypeValue = enumFlag[omnia.StreamType]{
	Default: omnia.VideoStreamType,
}

var streamTypeFlag = cli.GenericFlag{
	Name:    "type",
	Aliases: []string{"t"},
	Usage:   streamTypeValue.fmtUsage("target streamtype"),
	Value:   &streamTypeValue,
}

func App() *cli.App {
	return &cli.App{
		Name:   "omnia",
		Usage:  "some tools for nexxOmnia",
		Before: setupLogging,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "enable debug mode",
			},
			&cli.BoolFlag{
				Name:  "trace",
				Usage: "enable trace mode",
			},
		},
		Action: func(ctx *cli.Context) error {
			cli.ShowAppHelp(ctx)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "internal",
				Usage: "undocumented methods for internal use, will be removed",
				Subcommands: []*cli.Command{
					{
						Name:   "copy-cover-image",
						Usage:  "copies the name of the cover image of the id in the clipboard",
						Action: copyCoverCmd,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "input",
								Aliases: []string{"i"},
							},
						},
					},
					{
						Name:   "complete",
						Usage:  "complete the omnia db with local data",
						Action: completeCmd,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "config",
								Aliases: []string{"c"},
							},
							&cli.StringFlag{
								Name:    "input",
								Aliases: []string{"i"},
							},
						},
					},
					{
						Name:   "import",
						Usage:  "import data from pyton script",
						Action: importCmd,
					},
					{
						Name:   "merge",
						Usage:  "merge data into omnia",
						Action: mergeCmd,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "input",
								Aliases: []string{"i"},
							},
							&cli.StringFlag{
								Name:    "output",
								Aliases: []string{"o"},
							},
						},
					},
				},
			},
			{
				Name:  "publish",
				Usage: "publish a item",
				Flags: []cli.Flag{
					&configFileFlag,
					&streamTypeFlag,
				},
			},
		},
	}
}
