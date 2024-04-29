package tui

import (
	"strings"
	"tmarks/utils"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	bookmarks       []bookmark
	cursor          int
	keys            keyMap
	help            help.Model
	quittingMessage string
}

type errMsg struct {
	err error
}

type keyMap struct {
	Quit key.Binding
	Help key.Binding
	Up   key.Binding
	Down key.Binding
	Open key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Down, k.Up, k.Open, k.Quit, k.Help}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Down, k.Up, k.Open, k.Quit}, // first column
		{k.Help},                       // second column
	}
}

var DefaultKeyMap = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"), // actual keybindings
		key.WithHelp("q", "quit"),   // corresponding help text
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("j", "navigate down"),
	),
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("k", "navigate up"),
	),
	Open: key.NewBinding(
		key.WithKeys("enter", "return", "l"),
		key.WithHelp("enter", "open session"),
	),
}

func DisplayList() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		utils.HandleFatalError("display list", err)
	}
}

func initModel() model {
	return model{
		bookmarks: []bookmark{},
		cursor:    0,
		help:      help.New(),
		keys:      DefaultKeyMap,
	}
}

func (m model) Init() tea.Cmd {
	return getAllBookmarks
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Down):
			if m.cursor == len(m.bookmarks)-1 {
				m.cursor = 0
			} else {
				m.cursor++
			}
		case key.Matches(msg, DefaultKeyMap.Up):
			if m.cursor == 0 {
				m.cursor = len(m.bookmarks) - 1
			} else {
				m.cursor--
			}
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Open):
			sn := m.bookmarks[m.cursor].name
			return m, openTmuxSession(sn)
		}
	case bookmarksRetrievedMsg:
		m.bookmarks = msg.bookmarks
	case sessionOpenedMsg:
		m.quittingMessage = "\nLaunching the '" + msg.sessionName + "' session..."
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	if len(m.quittingMessage) > 0 {
		return m.quittingMessage + "\n\n"
	}
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString("Bookmarks\n")
	sb.WriteString("---------\n\n")
	if len(m.bookmarks) > 0 {
		for i, b := range m.bookmarks {
			var selector string
			if i == m.cursor {
				selector = "> "
			} else {
				selector = "  "
			}
			sb.WriteString(selector + b.name + "\n")
		}
	} else {
		sb.WriteString("No bookmarks saved...\n")
	}
	helpView := m.help.View(m.keys)
	sb.WriteString("\n" + helpView)
	return sb.String()
}
