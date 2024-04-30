package tmux

import (
	"os/exec"
)

func SessionIsUnderway(sn string) bool {
	sessionIsUserwayCmd := exec.Command("/bin/zsh", "-c", "tmux list-sessions | cut -d ':' -f 1 | grep '"+sn+"'")
	output, _ := sessionIsUserwayCmd.Output()
	return string(output) != ""
}

func OpenSession(sn string) error {
	cmd := exec.Command("tmux", "switch-client", "-t", sn)
	err := cmd.Start()
	if err != nil {
		return err
	} else {
		return nil
	}
}
