package cmd

import (
	"os"

	"github.com/android-project-46group/sgi-cli/api"
	"github.com/android-project-46group/sgi-cli/util"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sgi",
	Short: "About sakamichi group information",
	Long:  `SGI is a CLI libracy about sakamichi group information.`,
}

// Execute is the starting function when run cli
func Execute(config util.Config, api api.ApiCaller) {
	addVersionCmd(config.Version)
	addMemberCmd(api)
	addGroupCmd(api)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
