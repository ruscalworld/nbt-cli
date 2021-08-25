package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	InputFilePath  = ""
	OutputFilePath = ""
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
				Value:       InputFilePath,
				Destination: &OutputFilePath,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
