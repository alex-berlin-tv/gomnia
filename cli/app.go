package cli

import (
	"fmt"
	"strings"

	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/enums"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// Runs before each command. Used to set the logging level.
func onStartup(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}
	if ctx.Bool("trace") {
		log.SetLevel(log.TraceLevel)
	}
	return nil
}

type enumFlag[T ~string] struct {
	Default  enums.Enum[T]
	selected T
}

func (e *enumFlag[T]) Set(value string) error {
	rsl, err := enums.EnumByValue(e.Default, T(value))
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
	return fmt.Sprintf("[%s]", strings.Join(enums.EnumValues(e.Default), ", "))
}

func (e enumFlag[T]) fmtUsage(desc string) string {
	return fmt.Sprintf("%s %s", desc, e.fmtPossibleValues())
}

var configFileFlag = cli.PathFlag{
	Name:    "config",
	Aliases: []string{"c"},
	Usage:   "path to config file",
}

var idFlag = cli.StringSliceFlag{
	Name:    "id",
	Aliases: []string{"i"},
	Usage:   "target item(s) id",
}

var streamTypeValue = enumFlag[enums.StreamType]{Default: enums.VideoStreamType}

var streamTypeFlag = cli.GenericFlag{
	Name:    "type",
	Aliases: []string{"t"},
	Usage:   streamTypeValue.fmtUsage("target streamtype"),
	Value:   &streamTypeValue,
}

var ageValue = enumFlag[enums.AgeRestriction]{Default: enums.AgeRestriction0}

var ageFlag = cli.GenericFlag{
	Name:  "age",
	Usage: ageValue.fmtUsage("restrict to age"),
	Value: &ageValue,
}

var actionAfterRejectionValue = enumFlag[enums.ActionAfterRejection]{Default: enums.DeleteAfterRejection}

var actionAfterRejectionFlag = cli.GenericFlag{
	Name:  "action",
	Usage: actionAfterRejectionValue.fmtUsage("action after rejecting"),
	Value: &actionAfterRejectionValue,
}

func App() *cli.App {
	return &cli.App{
		Name:   "omnia",
		Usage:  "some tools for nexxOmnia",
		Before: onStartup,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "enable debug mode",
			},
			&cli.BoolFlag{
				Name:  "id-from-clipboard",
				Usage: "if no specific id is set via flag, the content of the clipboard will be used",
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
				Name:   "approve",
				Usage:  "approve an item",
				Action: approveCmd,
				Flags: []cli.Flag{
					&ageFlag,
					&idFlag,
					&configFileFlag,
					&streamTypeFlag,
					&cli.BoolFlag{
						Name:  "publish",
						Usage: "also publish the item",
					},
				},
			},
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
				Name:   "reject",
				Usage:  "reject an item",
				Action: unpublishCmd,
				Flags: []cli.Flag{
					&actionAfterRejectionFlag,
					&configFileFlag,
					&idFlag,
					&streamTypeFlag,
				},
			},
			{
				Name:   "publish",
				Usage:  "publish an item",
				Action: unpublishCmd,
				Flags: []cli.Flag{
					&configFileFlag,
					&idFlag,
					&streamTypeFlag,
				},
			},
			{
				Name:   "unblock",
				Usage:  "unblock a previously blocked item",
				Action: unblockCmd,
				Flags: []cli.Flag{
					&configFileFlag,
					&idFlag,
					&streamTypeFlag,
				},
			},
			{
				Name:   "unpublish",
				Usage:  "unpublish an item",
				Action: unpublishCmd,
				Flags: []cli.Flag{
					&configFileFlag,
					&idFlag,
					&streamTypeFlag,
					&cli.BoolFlag{
						Name:  "block",
						Usage: "publish will fail until item is unblocked",
					},
				},
			},
		},
	}
}
