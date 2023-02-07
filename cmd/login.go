package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/android-project-46group/sgi-cli/util"
	"github.com/spf13/cobra"
)

// Add login command to rootCmd
func addLoginCmd() {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "login",
		Long:  `For now, it just receives information from the user and stores it in a file.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Please provide your login information.\n\n")
			err := login()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(loginCmd)
}

// login is the main logic of the login function.
//
// Currently, it only reads the information entered by the user
// and saves it to the file.
func login() error {
	fmt.Printf(" BaseURL: ")
	url, err := readFromTerminal()
	if err != nil {
		return err
	}

	fmt.Printf(" APIKey: ")
	key, err := readFromTerminal()
	if err != nil {
		return err
	}

	ac := util.Account{
		BaseURL: url,
		APIKey:  key,
	}
	acJson, err := json.MarshalIndent(ac, "", "  ")
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Make the [util.CliDir] directory if needed.
	err = os.Mkdir(filepath.Join(home, util.CliDir), os.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	joined := filepath.Join(home, util.CliDir, util.AccountFile)
	if err = os.WriteFile(joined, acJson, 0644); err != nil {
		return err
	}
	fmt.Printf("\nLogin succeeded!\n")
	fmt.Printf("Your information is stored at: %s\n", joined)
	return nil
}

// readFromTerminal reads user's input from the terminal.
func readFromTerminal() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(input, "\n"), nil
}
