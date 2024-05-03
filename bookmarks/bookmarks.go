package bookmarks

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"tmarks/utils"

	"github.com/adrg/xdg"
)

type Bookmark struct {
	Id   string
	Name string
}

type appState struct {
	Bookmarks []Bookmark
}

func GetAll() []Bookmark {
	state, _ := readStateFile()
	return state.Bookmarks
}

func Add() {
	existingState, stateFilePath := readStateFile()
	b := createBookmarkForCurrentSession()
	addBookmarkToArray(&existingState, b)
	writeStateFile(stateFilePath, existingState)
}

func Delete(sn string) {
	state, stateFilePath := readStateFile()
	state.Bookmarks = slices.DeleteFunc(state.Bookmarks, func(e Bookmark) bool {
		return (e.Name == sn)
	})
	writeStateFile(stateFilePath, state)
}

func readStateFile() (appState, string) {
	filePath := fmt.Sprintf("%s/tmarks/app-state.json", xdg.StateHome)

	stat, fileExistsErr := os.Stat(filePath)
	if fileExistsErr != nil || stat.Size() == 0 {
		createEmptyFile(filePath)
	}

	b, readErr := os.ReadFile(filePath)
	if readErr != nil {
		utils.HandleFatalError("opening state file", readErr)
	}
	contents := unmarshalStateFile(b)
	return contents, filePath
}

func createEmptyFile(fp string) {
	state := appState{Bookmarks: []Bookmark{}}
	writeStateFile(fp, state)
}

func marshalStateFile(state appState) []byte {
	stateAsJson, marshalErr := json.Marshal(state)
	if marshalErr != nil {
		utils.HandleFatalError("marshalling state json", marshalErr)
	}
	return stateAsJson
}

func unmarshalStateFile(b []byte) appState {
	var contents appState
	unmarshalErr := json.Unmarshal(b, &contents)
	if unmarshalErr != nil {
		utils.HandleFatalError("unmarshalligg state json", unmarshalErr)
	}
	return contents
}

func createBookmarkForCurrentSession() Bookmark {
	cmd := exec.Command("tmux", "display-message", "-p", "#S")
	sessionName, err := cmd.Output()
	if err != nil {
		utils.HandleFatalError("obtaining tmux session name", err)
	}
	cleanSessionName := strings.ReplaceAll(string(sessionName), "\n", "")
	return Bookmark{Id: cleanSessionName, Name: cleanSessionName}
}

func addBookmarkToArray(as *appState, b Bookmark) {
	exists := slices.ContainsFunc((*as).Bookmarks, func(cb Bookmark) bool {
		return b.Id == cb.Id
	})
	if !exists {
		(*as).Bookmarks = append(as.Bookmarks, b)
	}
}

func writeStateFile(stateFilePath string, state appState) {
	json := marshalStateFile(state)
	writeErr := os.WriteFile(stateFilePath, []byte(json), 0666)
	if writeErr != nil {
		utils.HandleFatalError("writing state file", writeErr)
	}
}
