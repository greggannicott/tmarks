package main

import (
	"log"
	"os"

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
				Name:  "add",
				Usage: "Add a tmux session",
				Action: func(cCtx *cli.Context) error {
					addBookmark()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
