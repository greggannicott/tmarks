package tui

import (
	"tmarks/bookmarks"

	tea "github.com/charmbracelet/bubbletea"
)

type bookmark struct {
	name string
}

type bookmarksRetrievedMsg struct {
	bookmarks []bookmark
}

type bookmarkDeletedMsg struct {
	sessionName string
}

func getAllBookmarks() tea.Msg {
	var bks []bookmark
	for _, v := range bookmarks.GetAll() {
		bks = append(bks, bookmark{
			name: v.Name,
		})
	}
	return bookmarksRetrievedMsg{
		bookmarks: bks,
	}
}

func deleteBookmark(sn string) tea.Cmd {
	return func() tea.Msg {
		bookmarks.Delete(sn)
		return bookmarkDeletedMsg{sessionName: sn}
	}
}
