package main

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/hysmio/gcloud-interactive/gcloud"
	"github.com/hysmio/gcloud-interactive/gcloud/kubernetes"
	"github.com/manifoldco/promptui"
)

func main() {
	// Declare spinner - use https://github.com/briandowns/spinner to determine charset
	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	infoText := promptui.Styler(promptui.FGBold)
	extraText := promptui.Styler(promptui.FGFaint, promptui.FGBold)
	errorText := promptui.Styler(promptui.FGRed, promptui.FGBold)
	successText := promptui.Styler(promptui.FGGreen, promptui.FGBold)
	attrText := promptui.Styler(promptui.FGMagenta, promptui.FGBold)

	// Verify gcloud is installed
	spin.Prefix = "Verifying gcloud & kubectl are installed "
	spin.Start()
	gcloudbin, kubectlbin := gcloud.VerifyInstall()
	if gcloudbin != nil {
		spin.Stop()
		fmt.Printf("%s\n", errorText("gcloud is not installed"))

		os.Exit(1)
	}
	if kubectlbin != nil {
		spin.Stop()
		fmt.Printf("kubectl is not installed\n")
		prompt := promptui.Prompt{
			Label: "Would you like to install it? (Y/n) ",
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Couldn't get input")
			os.Exit(1)
		}
		if result != "n" {
			spin.Prefix = "Installing kubectl "
			spin.Start()
			_, err := gcloud.InstallComponent("kubectl")
			if err != nil {
				spin.Stop()
				fmt.Printf("Couldn't install kubectl %s\n", err)
				os.Exit(1)
			}
			spin.Stop()
			fmt.Printf("%s\n", successText("Installed kubectl"))
		} else {
			os.Exit(1)
		}
	}
	spin.Stop()
	fmt.Printf("%s\n", successText("gcloud & kubectl found!"))

	// TODO: Asyncronously read in more configuration details from gcloud

	spin.Prefix = "Getting clusters "
	spin.Start()
	clusters, err := kubernetes.GetClusters()
	if err != nil {
		spin.Stop()
		fmt.Printf("Couldn't get clusters: %s", err)
		os.Exit(1)
	}
	spin.Stop()

	cluster := kubernetes.Cluster{}

	if len(clusters) <= 0 {

	} else if len(clusters) == 1 {
		cluster = clusters[0]
	} else {
		prompt := promptui.Select{
			Label: "Select a Cluster",
			Items: clusters,
		}

		prompt.Run()
	}

	fmt.Printf("%s %s %s\n", infoText("Selected Cluster: "), attrText(cluster.Name), extraText("("+cluster.MasterIP+")"))

	projects, err := gcloud.GetAllProjects()

	fmt.Printf("Found project: %s %s\n", attrText(projects[0].Name), extraText("("+projects[0].ID+")"))

	// // Get the current project
	// spin.Prefix = "Getting current GCloud project: "
	// spin.Start()
	// project, err := gcloud.GetActiveProject()
	// if err != nil {
	// 	fmt.Printf("Couldn't get project: %s", err)
	// }
	// spin.Stop()
	// fmt.Printf("%s %s\n", successText("GCloud current project:"), attrText(project))

}

func renderMenu(initialProject gcloud.Project) {

}
