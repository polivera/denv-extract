/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/polivera/denv-extract/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var inputFilePath string

// fromFileCmd represents the fromFile command
var fromFileCmd = &cobra.Command{
	Use:   "fromFile",
	Short: "A brief description of your command",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("container name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		command(cmd, args)
	},
}

func init() {
	fromFileCmd.Flags().StringVarP(&inputFilePath, "file", "f", "", "--file /path/to/input/env")
	rootCmd.AddCommand(fromFileCmd)
}

func command(cmd *cobra.Command, args []string) {
	var (
		searchEnvVars []string
		resultEnvVars []string
		err           error
		serverName    string
		dumpPath      string
		containerInfo types.ContainerJSON
	)

	if inputFilePath == "" {
		fmt.Println("Missing environment file. You can pass it with -f /path/to/file")
		os.Exit(1)
	}

	if searchEnvVars, err = utils.ReadFromEnvFile(inputFilePath); err != nil {
		fmt.Println("[ERROR] - Can't get environment vars from file. " + err.Error())
		os.Exit(1)
	}

	if serverName, err = cmd.Flags().GetString("server"); err != nil {
		fmt.Println("[ERROR] - Can't get server flag. " + err.Error())
		os.Exit(1)
	}

	if containerInfo, err = utils.SearchSingleContainerInfo(args[0], serverName); err != nil {
		fmt.Println("[ERROR] - Can't get container info. " + err.Error())
		os.Exit(1)
	}

	// Split all the container vars
	splitContainerVars := map[string]string{}
	for _, containerVar := range containerInfo.Config.Env {
		splitVar := strings.Split(containerVar, "=")
		splitContainerVars[splitVar[0]] = splitVar[1]
	}

	// Split the search vars. For the empty ones set the container value
	for _, containerVar := range searchEnvVars {
		splitVar := strings.Split(containerVar, "=")
		if splitVar[1] == "" {
			resultEnvVars = append(
				resultEnvVars,
				fmt.Sprintf("%s=%s", splitVar[0], splitContainerVars[splitVar[0]]),
			)
		} else {
			resultEnvVars = append(
				resultEnvVars, fmt.Sprintf("%s=%s", splitVar[0], splitVar[1]),
			)
		}

	}

	// Write changes to file
	if dumpPath, err = utils.WriteToEnvFile(utils.CleanEnvArray(resultEnvVars)); err != nil {
		fmt.Println("[ERROR] - Can't dump result " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Result written to " + dumpPath)
}
