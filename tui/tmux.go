package tui

import (
	"tmarks/tmux"

	tea "github.com/charmbracelet/bubbletea"
)

type openSessionMsg struct{}

func openTmuxSession(sn string) tea.Msg {
	err := tmux.OpenSession(sn)
	if err != nil {
		return errMsg{err: err}
	}
	return openSessionMsg{}
}
