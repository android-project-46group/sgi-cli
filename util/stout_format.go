package util

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintTable(header []string, data [][]string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
	return nil
}
