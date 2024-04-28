package tui

import (
	"strings"
	"tmarks/utils"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	bookmarks []bookmark
	selected  int
	keys      keyMap
	help      help.Model
}

type keyMap struct {
	Quit key.Binding
	Help key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit}, // first column
		{k.Help}, // second column
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
		selected:  0,
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
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	case getBookmarksMsg:
		m.bookmarks = msg.bookmarks
	}
	return m, nil
}

func (m model) View() string {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString("Bookmarks\n")
	sb.WriteString("---------\n\n")
	if len(m.bookmarks) > 0 {
		for i, b := range m.bookmarks {
			var selector string
			if i == m.selected {
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
