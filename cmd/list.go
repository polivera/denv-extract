package cmd

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/polivera/denv-extract/helpers/docker"
	"github.com/polivera/denv-extract/helpers/survey"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list <container-name>",
	Short: "List environment variables of a container",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("container name is required")
		}
		return nil
	},
	Long: `Get a list of all environment variables from a container`,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		getEnvironmentVars(args[0], server)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func getEnvironmentVars(container string, server string) {
	var (
		err           error
		containerList []types.Container
		info          types.ContainerJSON
		fo            []string
	)

	if server == "" {
		server = "localhost"
	}

	dockerClient := docker.GetClient(server)
	if err = dockerClient.Connect(); err != nil {
		// deal with this error
	}
	if containerList, err = dockerClient.GetContainerByName(container); err != nil {
		// deal wit this
	}

	if len(containerList) == 0 {
		// container not found with name
		return
	} else if len(containerList) == 1 {
		info, _ = dockerClient.GetContainerInfo(containerList[0].ID)
	} else {
		var (
			options []string
			selCont int
		)

		for _, container := range containerList {
			options = append(options, getContainerName(container))
		}

		selCont, _, err = survey.GetAsk().Select(
			"Select container",
			options,
		)

		info, _ = dockerClient.GetContainerInfo(containerList[selCont].ID)
	}

	_, fo, err = survey.GetAsk().MultiSelect(
		"Select variables",
		info.Config.Env,
	)

	fmt.Println(fo)
}

func getContainerName(container types.Container) string {
	return container.Names[0][1:]
}
