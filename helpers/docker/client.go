package docker

import (
	"context"
	"github.com/docker/cli/cli/connhelper"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"net/http"
	"os"
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

	if dk.server != "localhost" {
		dk.cli, err = dk.connectToRemoteSSHServer()
	} else {
		dk.cli, err = client.NewClientWithOpts(client.FromEnv)
	}

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

// connectToRemoteSSHServer use ssh to use a remote docker server
// Thanks agbaraka(https://gist.github.com/agbaraka) for saving me days if not an eternity
// @see https://gist.github.com/agbaraka/654a218f8ea13b3da8a47d47595f5d05
func (dk *docker) connectToRemoteSSHServer() (*client.Client, error) {
	var (
		err    error
		helper *connhelper.ConnectionHelper
	)
	// Check if the url have the ssh
	if helper, err = connhelper.GetConnectionHelper(dk.server); err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		// No tls
		// No proxy
		Transport: &http.Transport{
			DialContext: helper.Dialer,
		},
	}
	var clientOpts []client.Opt

	clientOpts = append(clientOpts,
		client.WithHTTPClient(httpClient),
		client.WithHost(helper.Host),
		client.WithDialContext(helper.Dialer),
	)

	version := os.Getenv("DOCKER_API_VERSION")

	if version != "" {
		clientOpts = append(clientOpts, client.WithVersion(version))
	} else {
		clientOpts = append(clientOpts, client.WithAPIVersionNegotiation())
	}

	return client.NewClientWithOpts(clientOpts...)
}
