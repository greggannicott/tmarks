package utils

import (
	"fmt"
	"os"
)

var log *os.File

func SetLog(l *os.File) {
	log = l
}

func HandleFatalError(action string, err error) {
	message := fmt.Sprintf("Error %s: %s", action, err)
	fmt.Printf(message)
	log.WriteString(message + "\n")
	os.Exit(1)
}
