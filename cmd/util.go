package cmd

import (
	"fmt"
	"os"

	"github.com/android-project-46group/sgi-cli/util"
)

var (
	errMessage = `
BaseURL and APIKey are needed in your setting file.
If you are not logged in, please execute the following command.

$ sgi login
`
)

// validate checks before API call if the [util.Config] has the
// following items.
//
// - APIKey
// - BaseURL
//
// If it is not satisfied, print error message and exit with status 1.
func validate(cfg util.Config) {
	if cfg.Account.APIKey == "" || cfg.Account.BaseURL == "" {
		fmt.Println(errMessage)
		os.Exit(1)
	}
}
