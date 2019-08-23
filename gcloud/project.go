package gcloud

// Project GCloud representation of project
type Project struct {
	ID     string
	name   string
	number string
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
	CloudCommand("projects list")
}
