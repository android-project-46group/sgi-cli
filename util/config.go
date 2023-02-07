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
	// Path to Working directory related to CLI tools.
	// The path is described relative to the HOME directory.
	CliDir = ".sgi"

	// Name of the file where the account information is stored.
	AccountFile = "account.json"
)

// NewConfig returns a new [Config] instance.
// The configuration value is obtained from [AccountFile] under the ${HOME} directory.
// You need to pass the cli version as an argument.
func NewConfig(version string) (Config, error) {

	var cfg Config
	cfg.Version = version

	// Read account.json from a specific file.
	home, err := os.UserHomeDir()
	if err != nil {
		return cfg, err
	}
	joined := filepath.Join(home, CliDir, AccountFile)

	f, err := os.Open(joined)
	if err != nil {
		// If the error is [os.ErrNotExist], it's ok to continue without credential.
		if errors.Is(err, os.ErrNotExist) {
			return cfg, nil
		}
		return cfg, fmt.Errorf("Expected to read account information: %w\n", err)
	}
	defer f.Close()

	var ac Account
	if err := json.NewDecoder(f).Decode(&ac); err != nil {
		return cfg, err
	}
	cfg.Account = ac

	return cfg, nil
}
