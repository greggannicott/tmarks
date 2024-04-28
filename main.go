package main

import (
	"log"
	"os"
	"tmarks/bookmarks"
	"tmarks/tui"

	"github.com/urfave/cli/v2"
)

var Version = "Development"

func main() {
	app := &cli.App{
		Name:            "tmarks",
		Usage:           "Bookmarks for tmux",
		HideHelpCommand: true,
		Version:         Version,
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List and select from your bookmarks",
				Action: func(cCtx *cli.Context) error {
					tui.DisplayList()
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "Add a tmux session bookmark",
				Action: func(cCtx *cli.Context) error {
					bookmarks.Add()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
