/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/polivera/denv-extract/utils"
	"os"

	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Dump container env variables into file.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("container name is required")
		}
		return nil
	},
	Long: `Gather all environment variables from a container and dump them in a .env file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			dumpPath      string
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

		if dumpPath, err = utils.WriteToEnvFile(utils.CleanEnvArray(containerInfo.Config.Env)); err != nil {
			fmt.Println("[ERROR] - Can't dump result " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Result written to " + dumpPath)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
