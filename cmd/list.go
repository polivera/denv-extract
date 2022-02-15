package cmd

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/polivera/denv-extract/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
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
		var (
			varList       []string
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

		for _, envVar := range utils.CleanEnvArray(containerInfo.Config.Env) {
			varList = append(varList, strings.Split(envVar, "=")[0])
		}

		fmt.Println(strings.Join(varList, "\n"))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
