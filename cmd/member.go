package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/android-project-46group/sgi-cli/api"
	"github.com/android-project-46group/sgi-cli/util"
	"github.com/spf13/cobra"
)

// Add member command to rootCmd
func addMemberCmd(api api.ApiCaller) {
	memberCmd := &cobra.Command{
		Use:   "member",
		Short: "members information",
		Long:  `fetch and update members' information`,
	}

	listCmd := listMemberCmd(api)
	memberCmd.AddCommand(listCmd)

	rootCmd.AddCommand(memberCmd)
}

// listMemberCmd prints the all members of a specific groud.
func listMemberCmd(api api.ApiCaller) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "list information",
		Long:  `fetch all information`,
		Run: func(cmd *cobra.Command, args []string) {
			group, err := cmd.Flags().GetString("group")
			if err != nil {
				fmt.Println(err)
			}
			if group == "" {
				fmt.Println("ls command needs a group name")
				os.Exit(0)
			}

			members, err := api.ListMembers(group)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Print the members as a formatted table
			var data [][]string
			for _, member := range members {
				x := []string{
					strconv.Itoa(member.ID),
					member.Name,
				}
				data = append(data, x)
			}
			header := []string{
				"id",
				"name",
			}
			util.PrintTable(header, data)
		},
	}

	// -g nogizaka
	lsCmd.Flags().StringP("group", "g", "", "group name to fetch")

	return lsCmd
}
