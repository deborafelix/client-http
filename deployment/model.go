package deployment

import "time"

type Deployment struct {
	ID        string            `json:"id"`
	Labels    map[string]string `json:"labels"`
	Replicas  int               `json:"replicas"`
	Image     string            `json:"image"`
	Name      string            `json:"name"`
	Ports     []Port            `json:"ports"`
	CreatedAt time.Time         `json:"createdAt"`
}

type Port struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type DeploymentParams struct {
	ID       string            `json:"id"`
	Replicas int               `json:"replicas"`
	Image    string            `json:"image"`
	Labels   map[string]string `json:"labels"`
	Ports    []Port            `json:"ports"`
}
