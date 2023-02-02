package api

import "github.com/android-project-46group/sgi-cli/util"

// Api is a struct that implements [ApiCaller]
type Api struct {
	config util.Config
}

// New returns a new [ApiCaller] instance using passed configuration.
func New(config util.Config) ApiCaller {
	return &Api{
		config: config,
	}
}
