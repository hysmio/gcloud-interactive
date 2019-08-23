package gcloud

import (
	"os/exec"
)

// InstallComponent asks gcloud to install the component given in `component`
func InstallComponent(component string) (string, error) {
	gcloud := exec.Command("gcloud", "components", "install", component)
	output, err := gcloud.Output()
	return string(output), err
}
