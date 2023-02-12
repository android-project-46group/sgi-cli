package util_test

import (
	"github.com/android-project-46group/sgi-cli/util"
)

func ExamplePrintTable() {
	var header []string
	var data [][]string
	header = []string{
		"id",
		"name",
	}
	data = append(data, []string{
		"1",
		"John",
	})
	data = append(data, []string{
		"2",
		"Doe",
	})
	util.PrintTable(header, data)

	// Output:
	// +----+------+
	// | ID | NAME |
	// +----+------+
	// |  1 | John |
	// |  2 | Doe  |
	// +----+------+
}
