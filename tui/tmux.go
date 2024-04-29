package tui

import (
	"tmarks/tmux"

	tea "github.com/charmbracelet/bubbletea"
)

type sessionOpenedMsg struct{ sessionName string }

func openTmuxSession(sn string) tea.Cmd {
	return func() tea.Msg {
		err := tmux.OpenSession(sn)
		if err != nil {
			return errMsg{err: err}
		}
		return sessionOpenedMsg{sessionName: sn}
	}
}
