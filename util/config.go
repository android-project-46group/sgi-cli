package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Config contains the all settings for this cli.
// Create an instance of Config, by using NewConfig()
type Config struct {

	// Version is the application version.
	// Basically it consists of version and git-revision.
	Version string

	// Account contains the api settings of this cli
	Account Account
}

// Account contains the settings for api.
type Account struct {

	// BaseURL is the entrypoint of the api.
	BaseURL string

	// APIKey is needed to get information using API call
	APIKey string
}

var (
	// Path to an account setting file from HOME directory
	accountPath = ".sgi/account.json"
)

// NewConfig returns a new Config instance.
// The configuration value is obtained from accountPath under the ${HOME} directory.
// You need to pass the cli version as an argument.
func NewConfig(version string) (Config, error) {

	var cfg Config
	cfg.Version = version

	// Read account.json from a specific file.
	home, err := os.UserHomeDir()
	if err != nil {
		return cfg, err
	}
	joined := filepath.Join(home, accountPath)

	f, err := os.Open(joined)
	if err != nil {
		return cfg, errors.New(fmt.Sprintf("Expected to read account information from: %s\n", joined))
	}
	defer f.Close()

	var ac Account
	if err := json.NewDecoder(f).Decode(&ac); err != nil {
		return cfg, err
	}
	cfg.Account = ac

	return cfg, nil
}
