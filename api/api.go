package api

import "github.com/android-project-46group/sgi-cli/util"

// struct that implements ApiCaller
type Api struct {
	config util.Config
}

func New(config util.Config) ApiCaller {
	return &Api{
		config: config,
	}
}
