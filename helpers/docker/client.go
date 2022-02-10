package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type Client interface {
	Connect() error
	GetContainerByName(name string) ([]types.Container, error)
	GetContainerInfo(containerId string) (types.ContainerJSON, error)
}

type docker struct {
	server string
	cli    *client.Client
}

// GetClient return an instance of the Client interface
func GetClient(server string) Client {
	if server == "" {
		server = "localhost"
	}
	dockerClient := docker{server: server}
	return &dockerClient
}

// Connect client to a docker instance on docker.server
func (dk *docker) Connect() error {
	var err error

	dk.cli, err = client.NewClientWithOpts(client.FromEnv)
	return err
}

// GetContainerByName return a list of containers that match the name
// The criteria will match any container when its name is equal or start with
// the given parameter
func (dk *docker) GetContainerByName(name string) ([]types.Container, error) {
	listFilters := filters.NewArgs(filters.Arg("name", name))

	return dk.cli.ContainerList(
		context.Background(),
		types.ContainerListOptions{Filters: listFilters},
	)
}

// GetContainerInfo get inspect result of a given container
func (dk *docker) GetContainerInfo(containerId string) (types.ContainerJSON, error) {
	return dk.cli.ContainerInspect(
		context.Background(),
		containerId,
	)
}
