package tui

import (
	"tmarks/bookmarks"

	tea "github.com/charmbracelet/bubbletea"
)

type bookmark struct {
	id   string
	name string
}

type bookmarksRetrievedMsg struct {
	bookmarks []bookmark
}

func getAllBookmarks() tea.Msg {
	var bks []bookmark
	for _, v := range bookmarks.GetAll() {
		bks = append(bks, bookmark{
			id:   v.Id,
			name: v.Name,
		})
	}
	return bookmarksRetrievedMsg{
		bookmarks: bks,
	}
}
