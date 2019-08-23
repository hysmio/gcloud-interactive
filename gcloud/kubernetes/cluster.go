package kubernetes

import (
	"strings"

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
	// output := "NAME              LOCATION                     MASTER_VERSION             MASTER_IP           MACHINE_TYPE          NODE_VERSION         NUM_NODES         STATUS\n"
	// output += "automated-cloud   australia-south-east-a       some-kubernetes-version    192.168.213.123     A v smol one          node version         2                 RUNNING"

	clusters := make([]Cluster, 0)

	location := strings.Index(output, "LOCATION")
	masterVersion := strings.Index(output, "MASTER_VERSION")
	masterIP := strings.Index(output, "MASTER_IP")
	machineType := strings.Index(output, "MACHINE_TYPE")
	nodeVersion := strings.Index(output, "NODE_VERSION")
	numNodes := strings.Index(output, "NUM_NODES")
	status := strings.Index(output, "STATUS")

	for pos, value := range strings.Split(output, "\n") {
		if pos == 0 {
			continue
		}

		if len(value) > status {
			cluster := Cluster{
				Name:          strings.TrimSpace(value[0:location]),
				Location:      strings.TrimSpace(value[location:masterVersion]),
				MasterVersion: strings.TrimSpace(value[masterVersion:masterIP]),
				MasterIP:      strings.TrimSpace(value[masterIP:machineType]),
				MachineType:   strings.TrimSpace(value[machineType:nodeVersion]),
				NodeVersion:   strings.TrimSpace(value[nodeVersion:numNodes]),
				NumNodes:      strings.TrimSpace(value[numNodes:status]),
				Status:        strings.TrimSpace(value[status:len(value)]),
			}

			clusters = append(clusters, cluster)
		}
	}

	return clusters, err
}
