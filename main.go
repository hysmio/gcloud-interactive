package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/hysmio/gcloud-interactive/gcloud"
	"github.com/manifoldco/promptui"
)

func main() {
	// Declare spinner - use https://github.com/briandowns/spinner to determine charset
	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	infoText := promptui.Styler(promptui.FGBold)
	// extraText := promptui.Styler(promptui.FGFaint, promptui.FGBold)
	errorText := promptui.Styler(promptui.FGRed, promptui.FGBold)
	successText := promptui.Styler(promptui.FGGreen, promptui.FGBold)
	// attrText := promptui.Styler(promptui.FGMagenta, promptui.FGBold)

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

	// spin.Prefix = "Getting clusters "
	// spin.Start()
	// clusters, err := kubernetes.GetClusters()
	// if err != nil {
	// 	spin.Stop()
	// 	fmt.Printf("Couldn't get clusters: %s", err)
	// 	os.Exit(1)
	// }
	// spin.Stop()

	// cluster := kubernetes.Cluster{}

	// if len(clusters) <= 0 {

	// } else if len(clusters) == 1 {
	// 	cluster = clusters[0]
	// } else {
	// 	prompt := promptui.Select{
	// 		Label: "Select a Cluster",
	// 		Items: clusters,
	// 	}

	// 	prompt.Run()
	// }

	// fmt.Printf("%s %s %s\n", infoText("Selected Cluster: "), attrText(cluster.Name), extraText("("+cluster.MasterIP+")"))

	spin.Prefix = infoText("Getting projects ")
	spin.Start()
	projects, _ := gcloud.GetAllProjects()
	spin.Stop()

	project := gcloud.Project{}

	if len(projects) == 0 {
		// TODO: Add ability to create projects
		println(errorText("You have no projects, exiting..."))
		os.Exit(1)
	} else if len(projects) == 1 {
		project = projects[0]
		println(infoText("Found 1 project!"))
		activeProject, _ := gcloud.GetActiveProject()

		println(activeProject)

		if strings.TrimSpace(activeProject) != project.ID {
			println(errorText("We only found one project, but it is not your currently active project"))
			prompt := promptui.Prompt{
				Label: infoText("Would you like to set ") +
					project.Format() +
					infoText(" as your current project? (Y/n) "),
			}

			result, _ := prompt.Run()

			if result != "n" {
				spin.Prefix = infoText("Setting current project to " + project.Format())
				spin.Start()
				gcloud.SetActiveProject(project.ID)
				spin.Stop()
				println(successText("Set ") + project.Format() + successText(" as active project"))
			}
		}
	} else {
		// TODO: Get current project, let user choose between projects
	}

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
