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

			d, err := cmd.Flags().GetBool("data")
			if err != nil {
				fmt.Println(err)
			}

			members, err := api.ListMembers(group)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			printMemberList(members, d)
		},
	}

	// -g nogizaka
	lsCmd.Flags().StringP("group", "g", "", "group name to fetch")
	// -d, --data
	lsCmd.Flags().BoolP("data", "d", false, "print all data")

	return lsCmd
}

// Print the members as a formatted table
func printMemberList(members []api.Member, all bool) {

	var header []string
	var data [][]string

	if all {
		header = []string{
			"id",
			"name",
			"birthday",
			"height",
			"blood",
			"generation",
			"blog",
			"img",
		}
		for _, member := range members {
			x := []string{
				strconv.Itoa(member.ID),
				member.Name,
				member.Birthday,
				member.Height,
				member.Blood,
				member.Generation,
				member.BlogURL,
				member.ImgURL,
			}
			data = append(data, x)
		}
	} else {
		header = []string{
			"id",
			"name",
			"generation",
		}
		for _, member := range members {
			x := []string{
				strconv.Itoa(member.ID),
				member.Name,
				member.Generation,
			}
			data = append(data, x)
		}

	}

	util.PrintTable(header, data)
}
