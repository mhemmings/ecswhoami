// Package ecswhoami does a task metadata lookup from inside AWS Elastic Container Service.
// For more details see: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-metadata-endpoint-v3.html
package ecswhoami

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-metadata-endpoint-v3.html
type Metadata struct {
	DockerID      string            `json:"DockerId"`
	Name          string            `json:"Name"`
	DockerName    string            `json:"DockerName"`
	Image         string            `json:"Image"`
	ImageID       string            `json:"ImageID"`
	Labels        map[string]string `json:"Labels"`
	DesiredStatus string            `json:"DesiredStatus"`
	KnownStatus   string            `json:"KnownStatus"`
	Limits        Limits            `json:"Limits"`
	CreatedAt     time.Time         `json:"CreatedAt"`
	StartedAt     time.Time         `json:"StartedAt"`
	Type          string            `json:"Type"`
	Networks      []Network         `json:"Networks"`
}

type Limits struct {
	CPU    int `json:"CPU"`
	Memory int `json:"Memory"`
}

type Network struct {
	NetworkMode   string   `json:"NetworkMode"`
	IPv4Addresses []string `json:"IPv4Addresses"`
}

// GetLabel returns the label for the given name, or an empty string if it's not set.
func (m Metadata) GetLabel(name string) string {
	l, _ := m.Labels[name]
	return l
}

// Cluster returns the com.amazonaws.ecs.cluster, or an empty string if it's not set.
func (m Metadata) Cluster() string {
	return m.GetLabel("com.amazonaws.ecs.cluster")
}

// ContainerName returns the com.amazonaws.ecs.container-name, or an empty string if it's not set.
func (m Metadata) ContainerName() string {
	return m.GetLabel("com.amazonaws.ecs.container-name")
}

// TaskArn returns the com.amazonaws.ecs.task-arn, or an empty string if it's not set.
func (m Metadata) TaskArn() string {
	return m.GetLabel("com.amazonaws.ecs.task-arn")
}

// TaskDefinitionFamily returns the com.amazonaws.ecs.task-definition-family, or an empty string if it's not set.
func (m Metadata) TaskDefinitionFamily() string {
	return m.GetLabel("com.amazonaws.ecs.task-definition-family")
}

// TaskDefinitionVersion returns the com.amazonaws.ecs.task-definition-version, or an empty string if it's not set.
func (m Metadata) TaskDefinitionVersion() string {
	return m.GetLabel("com.amazonaws.ecs.task-definition-version")
}

const envVar = "ECS_CONTAINER_METADATA_URI"

var (
	// ErrEnvNotSet defines an error for when the ECS URI environment variable is not set.
	ErrEnvNotSet = fmt.Errorf("%s not set. Are you in an ECS environment?", envVar)

	// A http.Client with a timeout.
	httpClient = &http.Client{
		Timeout: time.Second * 5,
	}
)

// Lookup calls the URL ECS_CONTAINER_METADATA_URI, returning the response as a Metadata.
func Lookup() (Metadata, error) {
	uri := os.Getenv(envVar)
	if uri == "" {
		return Metadata{}, ErrEnvNotSet
	}
	res, err := httpClient.Get(uri)
	if err != nil {
		return Metadata{}, err
	}
	defer res.Body.Close()
	var meta Metadata
	err = json.NewDecoder(res.Body).Decode(&meta)
	if err != nil {
		return Metadata{}, err
	}
	return meta, nil
}
