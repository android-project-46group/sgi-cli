package api

import (
	"net/http"

	"github.com/android-project-46group/sgi-cli/util"
)

// Api is a struct that implements [ApiCaller]
type Api struct {

	// config contains the all settings for api.
	config util.Config

	// client is used for HTTP communication.
	client http.Client
}

// New returns a new [ApiCaller] instance using passed configuration.
func New(config util.Config) ApiCaller {
	client := NewClient()
	return &Api{
		config: config,
		client: client,
	}
}
