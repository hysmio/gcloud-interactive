package kubernetes

import (
	"github.com/hysmio/gcloud-interactive/parser"

	"github.com/hysmio/gcloud-interactive/gcloud"
)

// Cluster for representing a Kubernetes Cluster
type Cluster struct {
	Name          string
	Location      string
	MasterVersion string
	MasterIP      string
	MachineType   string
	NodeVersion   string
	NumNodes      string
	Status        string
}

// CreateCluster creates a cluster on GCloud
// TODO: Add Cluster struct & return it
func CreateCluster(name string) error {
	return gcloud.CloudCommand("container", "clusters", "create", name).Run()
}

// GetClusters gets all clusters under the current project
func GetClusters() ([]Cluster, error) {
	raw, err := gcloud.CloudCommand("container", "clusters", "list").Output()
	output := string(raw)

	clusters := make([]Cluster, 0)
	rows := parser.ParseTable(output)

	for _, r := range rows {
		cluster := Cluster{
			Name:          r["NAME"],
			Location:      r["LOCATION"],
			MasterVersion: r["MASTER_VERSION"],
			MasterIP:      r["MASTER_IP"],
			MachineType:   r["MACHINE_TYPE"],
			NodeVersion:   r["NODE_VERSION"],
			NumNodes:      r["NUM_NODES"],
			Status:        r["STATUS"],
		}

		clusters = append(clusters, cluster)
	}

	return clusters, err
}
