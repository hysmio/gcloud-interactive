package gcloud

import (
	"github.com/hysmio/gcloud-interactive/parser"
)

// Project GCloud representation of project
type Project struct {
	ID     string
	Name   string
	Number string
}

// GetActiveProject gets the currently active project using `gcloud config get-value project`
func GetActiveProject() (string, error) {
	return GetConfig("project")
}

// SetActiveProject sets the current project to projectId
func SetActiveProject(projectID string) error {
	return SetConfig("project", projectID)
}

// GetAllProjects gets all the projects GCloud has
func GetAllProjects() ([]Project, error) {
	raw, err := CloudCommand("projects", "list").Output()

	projects := make([]Project, 0)

	if err == nil {
		rawValues := parser.ParseTable(string(raw))

		for _, value := range rawValues {
			project := Project{
				ID:     value["PROJECT_ID"],
				Name:   value["NAME"],
				Number: value["PROJECT_NUMBER"],
			}

			projects = append(projects, project)
		}
	}

	return projects, err
}
