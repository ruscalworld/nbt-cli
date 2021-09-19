package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var (
	InputFilePath  = ""
	OutputFilePath = ""
	ArrayItemLimit = 50
)

func main() {
	app := cli.App{
		Name:        "NBT Browser",
		Description: "Simple command-line tool for working with Minecraft's NBT format",
		Commands: []*cli.Command{
			{
				Name:   "tree",
				Action: FileAction(PrintTree),
			},
			{
				Name:   "set",
				Action: FileAction(SetValue),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "key",
						Aliases:  []string{"k"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "value",
						Aliases:  []string{"v"},
						Required: true,
					},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"f"},
				Required:    true,
				Destination: &InputFilePath,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Destination: &OutputFilePath,
			},
			&cli.IntFlag{
				Name:        "limit",
				Aliases:     []string{"l"},
				Value:       50,
				Destination: &ArrayItemLimit,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
		return
	}
}
