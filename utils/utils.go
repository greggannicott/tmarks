package utils

import (
	"fmt"
	"os"
)

func HandleFatalError(action string, err error) {
	fmt.Printf("Error %s: %s", action, err)
	os.Exit(1)
}
