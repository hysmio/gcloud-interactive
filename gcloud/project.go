package gcloud

import (
	"os/exec"
)

func getConfig(key string) (string, error) {
	cmd := exec.Command("gcloud", "config", "get-value", key)
	output, err := cmd.Output()

	return string(output), err
}

// GetActiveProject gets the currently active project using `gcloud config get-value project`
func GetActiveProject() (string, error) {
	return getConfig("project")
}
