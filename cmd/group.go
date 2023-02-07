package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/android-project-46group/sgi-cli/api"
	"github.com/android-project-46group/sgi-cli/util"
	"github.com/spf13/cobra"
)

// Add member command to rootCmd
func addGroupCmd(api api.ApiCaller, cfg util.Config) {
	// groupCmd represents the group command
	var groupCmd = &cobra.Command{
		Use:   "group",
		Short: "groups information",
		Long:  `fetch and update groups' information`,
	}
	listCmd := listGroupCmd(api, cfg)
	groupCmd.AddCommand(listCmd)

	rootCmd.AddCommand(groupCmd)
}

// listGroupCmd prints the all groups.
func listGroupCmd(api api.ApiCaller, cfg util.Config) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "list information",
		Long:  `fetch all information`,
		Run: func(cmd *cobra.Command, args []string) {

			validate(cfg)

			j, err := cmd.Flags().GetBool("json")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			groups, err := api.ListGroups()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if j {
				printGroupsListJson(groups)
			} else {
				printGroupsList(groups)
			}
		},
	}

	// -j, --json
	lsCmd.Flags().BoolP("json", "j", false, "print as a json format")

	return lsCmd
}

// Print the groups as a formatted table
func printGroupsList(groups []api.Group) {

	var header []string
	var data [][]string

	header = []string{
		"id",
		"name",
	}
	for _, group := range groups {
		x := []string{
			strconv.Itoa(group.ID),
			group.Name,
		}
		data = append(data, x)
	}

	util.PrintTable(header, data)
}

// Print the groups as a json format
func printGroupsListJson(groups []api.Group) {

	b, err := json.Marshal(groups)
	if err != nil {
		fmt.Println("failed to marshal to json: %w", err)
		os.Exit(1)
	}

	// Final output
	fmt.Printf(string(b))
}
