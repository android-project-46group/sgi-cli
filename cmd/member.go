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
func addMemberCmd(api api.ApiCaller, cfg util.Config) {
	memberCmd := &cobra.Command{
		Use:   "member",
		Short: "members information",
		Long:  `fetch and update members' information`,
	}

	listCmd := listMemberCmd(api, cfg)
	memberCmd.AddCommand(listCmd)

	rootCmd.AddCommand(memberCmd)
}

type listMember struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Generation string `json:"generation"`
}

type listMemberAll struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Birthday   string `json:"birthday"`
	Height     string `json:"height"`
	Blood      string `json:"blood"`
	Generation string `json:"generation"`
	BlogURL    string `json:"blog"`
	ImgURL     string `json:"img"`
}

// listMemberCmd prints the all members of a specific groud.
func listMemberCmd(api api.ApiCaller, cfg util.Config) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "list information",
		Long:  `fetch all information`,
		Run: func(cmd *cobra.Command, args []string) {

			validate(cfg)

			group, err := cmd.Flags().GetString("group")
			if err != nil {
				fmt.Println(err)
			}

			d, err := cmd.Flags().GetBool("data")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			j, err := cmd.Flags().GetBool("json")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			members, err := api.ListMembers(group)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if j {
				printMemberListJson(members, d)
			} else {
				printMemberList(members, d)
			}
		},
	}

	// -g nogizaka
	lsCmd.Flags().StringP("group", "g", "", "group name to fetch")
	// -d, --data
	lsCmd.Flags().BoolP("data", "d", false, "print all data")
	// -j, --json
	lsCmd.Flags().BoolP("json", "j", false, "print as a json format")

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

// Print the members as a json format
func printMemberListJson(members []api.Member, all bool) {

	if all {
		var ms []listMemberAll
		for _, member := range members {
			m := listMemberAll{
				ID:         member.ID,
				Name:       member.Name,
				Birthday:   member.Birthday,
				Height:     member.Height,
				Blood:      member.Blood,
				Generation: member.Generation,
				BlogURL:    member.BlogURL,
				ImgURL:     member.ImgURL,
			}
			ms = append(ms, m)
		}

		b, err := json.Marshal(ms)
		if err != nil {
			fmt.Println("failed to marshal to json: %w", err)
			os.Exit(1)
		}

		// Final output
		fmt.Printf(string(b))
	} else {
		var ms []listMember

		for _, member := range members {
			m := listMember{
				ID:         member.ID,
				Name:       member.Name,
				Generation: member.Generation,
			}
			ms = append(ms, m)
		}

		b, err := json.Marshal(ms)
		if err != nil {
			fmt.Println("failed to marshal to json: %w", err)
			os.Exit(1)
		}

		// Final output
		fmt.Printf(string(b))
	}
}
