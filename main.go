package main

import (
	"fmt"

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
	}
	cmd.Execute(config)
}
