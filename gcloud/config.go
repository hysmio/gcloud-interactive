package gcloud

import (
	"os/exec"
	"strings"
)

// GetConfig gets a configuration value from the given key
func GetConfig(key string) (string, error) {
	cmd := exec.Command("gcloud", "config", "get-value", key)
	output, err := cmd.Output()

	return strings.TrimSpace(string(output)), err
}

// SetConfig sets a configuration value with the given key
func SetConfig(key string, value string) error {
	cmd := exec.Command("gcloud", "config", "set", key, value)
	err := cmd.Run()

	return err
}
