package gcloud

import (
	"os/exec"
)

// CloudCommand sends a command to gcloud
func CloudCommand(arguments ...string) *exec.Cmd {
	return exec.Command("gcloud", arguments...)
}
