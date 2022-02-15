package utils

import (
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/polivera/denv-extract/helpers/docker"
	"github.com/polivera/denv-extract/helpers/survey"
)

func SearchSingleContainerInfo(containerCriteria string, server string) (types.ContainerJSON, error) {
	var (
		err           error
		containerList []types.Container
		info          types.ContainerJSON
		containerId   string
	)

	dockerClient := docker.GetClient(server)
	if err = dockerClient.Connect(); err != nil {
		return info, err
	}
	if containerList, err = dockerClient.GetContainerByName(containerCriteria); err != nil {
		return info, err
	}

	switch len(containerList) {
	case 0:
		return info, nil
	case 1:
		containerId = containerList[0].ID
	default:
		if containerId, err = selectFromContainerList(containerList); err != nil {
			return info, err
		}
	}

	if info, err = dockerClient.GetContainerInfo(containerId); err != nil {
		return info, err
	}

	if info.Config == nil {
		return info, errors.New("no container found")
	}

	return info, err
}

// selectFromContainerList pops a list of the found containers and return the container ID
func selectFromContainerList(containerList []types.Container) (string, error) {
	var (
		options  []string
		selIndex int
		err      error
	)

	for _, container := range containerList {
		options = append(options, getContainerName(container))
	}

	if selIndex, _, err = survey.GetAsk().Select("Select container", options); err != nil {
		return "", err
	}

	return containerList[selIndex].ID, err
}

// getContainerName return the container name from container.Names
func getContainerName(container types.Container) string {
	return container.Names[0][1:]
}
