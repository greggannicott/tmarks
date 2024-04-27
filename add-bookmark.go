package main

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

type bookmark struct {
	Id   string
	Name string
}

type appState struct {
	Bookmarks []bookmark
}

func addBookmark() {
	existingState, stateFilePath := readStateFile()
	b := createBookmarkForCurrentSession()
	addBookmarkToArray(&existingState, b)
	stateAsJson := marshalStateFile(existingState)
	writeStateFile(stateFilePath, stateAsJson)
}

func readStateFile() (appState, string) {
	filePath := fmt.Sprintf("%s/tmarks/app-state.json", xdg.StateHome)

	_, fileExistsErr := os.Stat(filePath)
	if fileExistsErr != nil {
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
	state := appState{Bookmarks: []bookmark{}}
	stateAsJson := marshalStateFile(state)
	writeStateFile(fp, stateAsJson)
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

func createBookmarkForCurrentSession() bookmark {
	cmd := exec.Command("tmux", "display-message", "-p", "#S")
	sessionName, err := cmd.Output()
	if err != nil {
		utils.HandleFatalError("obtaining tmux session name", err)
	}
	cleanSessionName := strings.ReplaceAll(string(sessionName), "\n", "")
	return bookmark{Id: cleanSessionName, Name: cleanSessionName}
}

func addBookmarkToArray(as *appState, b bookmark) {
	exists := slices.ContainsFunc((*as).Bookmarks, func(cb bookmark) bool {
		return b.Id == cb.Id
	})
	if !exists {
		(*as).Bookmarks = append(as.Bookmarks, b)
	}
}

func writeStateFile(stateFilePath string, json []byte) {
	writeErr := os.WriteFile(stateFilePath, []byte(json), 0666)
	if writeErr != nil {
		utils.HandleFatalError("writing state file", writeErr)
	}
}
