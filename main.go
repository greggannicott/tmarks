package main

import (
	"fmt"
	"log"
	"os"
	"tmarks/bookmarks"
	"tmarks/tui"
	"tmarks/utils"

	"github.com/adrg/xdg"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

var Version = "Development"

func main() {
	l := getLogFile()
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
					tui.DisplayList(l)
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
	l.Close()
}

func getLogFile() *os.File {
	appDirPath := fmt.Sprintf("%s/tmarks", xdg.DataHome)

	_, appDirStatErr := os.Stat(appDirPath)
	if appDirStatErr != nil {
		os.MkdirAll(appDirPath, 0777)
	}

	logPath := fmt.Sprintf("%s/tmarks.log", appDirPath)
	l, err := tea.LogToFile(logPath, "debug")
	if err != nil {
		utils.HandleFatalError("creating log file", err)
	}
	return l
}
