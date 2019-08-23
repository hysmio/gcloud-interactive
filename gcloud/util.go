package gcloud

import (
	"os/exec"
)

// VerifyInstall verifies that the `gcloud` binary is installed
func VerifyInstall() (error, error) {
	gcloud := exec.Command("gcloud", "version")
	kubectl := exec.Command("kubectl")
	return gcloud.Run(), kubectl.Run()
}
