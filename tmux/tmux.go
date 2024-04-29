package tmux

import (
	"os/exec"
)

func OpenSession(sn string) error {
	cmd := exec.Command("tmux", "switch-client", "-t", sn)
	err := cmd.Start()
	if err != nil {
		return err
	} else {
		return nil
	}
}
