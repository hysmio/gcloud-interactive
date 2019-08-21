package main

import (
	"fmt"
	// "github.com/manifoldco/promptui"
	"github.com/hysmio/gcloud-interactive/gcloud"
)

func main() {
	project, err := gcloud.GetActiveProject()
	if err != nil {
		fmt.Printf("Couldn't get project: %s", err)
	}

	fmt.Printf("Current Project: %s", project)
}
