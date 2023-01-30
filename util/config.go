package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Version string
	Account Account
}

type Account struct {
	BaseURL string
	APIKey  string
}

var (
	// Path to an account setting file from HOME directory
	accountPath = ".sgi/account.json"
)

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
