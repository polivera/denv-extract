/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/polivera/denv-extract/helpers/survey"
	"github.com/polivera/denv-extract/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Dump selected container env variables into file.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("container name is required")
		}
		return nil
	},
	Long: `Gather selected environment variables from a container and dump them in a .env file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			dumpPath      string
			varList       []string
			cleanedList   []string
			indexList     []int
			server        string
			err           error
			containerInfo types.ContainerJSON
		)

		if server, err = cmd.Flags().GetString("server"); err != nil {
			fmt.Println("[ERROR] - Can't get server flag. " + err.Error())
			os.Exit(1)
		}

		if containerInfo, err = utils.SearchSingleContainerInfo(args[0], server); err != nil {
			fmt.Println("[ERROR] - Can't get container info. " + err.Error())
			os.Exit(1)
		}

		cleanedList = utils.CleanEnvArray(containerInfo.Config.Env)
		for _, envVar := range cleanedList {
			varList = append(varList, strings.Split(envVar, "=")[0])
		}

		indexList, _, _ = survey.GetAsk().MultiSelect("Select what you want to save", varList)
		varList = nil //clean this, so I don't have to declare a new one

		for _, i := range indexList {
			varList = append(varList, cleanedList[i])
		}

		if dumpPath, err = utils.WriteToEnvFile(varList); err != nil {
			fmt.Println("[ERROR] - Can't dump result " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Result written to " + dumpPath)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
