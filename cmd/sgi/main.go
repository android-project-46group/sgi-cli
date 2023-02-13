package main

import (
	"fmt"
	"log"

	"github.com/android-project-46group/sgi-cli/api"
	"github.com/android-project-46group/sgi-cli/cmd"
	"github.com/android-project-46group/sgi-cli/util"
)

var (
	version  = "0.1.0"
	revision = "HEAD"
)

func main() {
	ver := fmt.Sprintf("%s - %s", version, revision)
	config, err := util.NewConfig(ver)
	if err != nil {
		log.Fatal(err)
	}
	api := api.New(config)

	cmd.Execute(config, api)
}
