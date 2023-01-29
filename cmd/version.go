package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Add version command to rootCmd
func addVersionCmd(version string) {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of sgi-tool",
		Long:  `Print the version of sgi-tool.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	rootCmd.AddCommand(versionCmd)
}
