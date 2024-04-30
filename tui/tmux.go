package tui

import (
	"tmarks/tmux"

	tea "github.com/charmbracelet/bubbletea"
)

type sessionOpenedMsg struct{ sessionName string }
type sessionNotFound struct{ sessionName string }

func openTmuxSession(sn string) tea.Cmd {
	return func() tea.Msg {
		sessionFound := tmux.SessionIsUnderway(sn)
		if sessionFound {
			err := tmux.OpenSession(sn)
			if err != nil {
				return errMsg{err: err}
			}
			return sessionOpenedMsg{sessionName: sn}
		} else {
			return sessionNotFound{sessionName: sn}
		}
	}
}
