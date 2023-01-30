package main

import (
	"fmt"

	"github.com/android-project-46group/sgi-cli/api"
	"github.com/android-project-46group/sgi-cli/cmd"
	"github.com/android-project-46group/sgi-cli/util"
)

const version = "0.1.0"

var (
	revision = "HEAD"
)

func main() {
	config := util.Config{
		Version: fmt.Sprintf("%s - %s", version, revision),
		BaseURL: "https://kokoichi0206.mydns.jp/api/v1",
		APIKey: "my_api_key",
	}
	api := api.New(config)

	cmd.Execute(config, api)
}
