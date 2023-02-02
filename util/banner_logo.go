package util

import (
	"bytes"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

// Logo format of SGI.
const CLI_LOGO = `
{{ .AnsiColor.Black }}██████████████████████████████████████████████████████████████████████
███████████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}████████{{ .AnsiColor.Green }}█████{{ .AnsiColor.Black }}███████████{{ .AnsiColor.Green }}█████{{ .AnsiColor.Black }}█████████{{ .AnsiColor.Green }}██████{{ .AnsiColor.Black }}█████████
██████████████{{ .AnsiColor.Green }}███{{ .AnsiColor.Black }}██████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}█████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}███{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}██████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}███████████
████████████{{ .AnsiColor.Green }}█████{{ .AnsiColor.Black }}██████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}█████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}███████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}█████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}█████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}███████████
███████████{{ .AnsiColor.Green }}██████{{ .AnsiColor.Black }}████████{{ .AnsiColor.Green }}███{{ .AnsiColor.Black }}███████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}████████████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}███████████
█████████{{ .AnsiColor.Green }}████████{{ .AnsiColor.Black }}██████████{{ .AnsiColor.Green }}███{{ .AnsiColor.Black }}█████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}████{{ .AnsiColor.Green }}███{{ .AnsiColor.Black }}█████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}███████████
████████{{ .AnsiColor.Green }}█████████{{ .AnsiColor.Black }}█████████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}███████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}█████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}█████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}███████████
██████{{ .AnsiColor.Green }}███████████{{ .AnsiColor.Black }}██████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}█████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}█████████{{ .AnsiColor.Green }}██{{ .AnsiColor.Black }}███████████
█████{{ .AnsiColor.Green }}████████████{{ .AnsiColor.Black }}████████{{ .AnsiColor.Green }}█████{{ .AnsiColor.Black }}███████████{{ .AnsiColor.Green }}███████{{ .AnsiColor.Black }}███████{{ .AnsiColor.Green }}██████{{ .AnsiColor.Black }}█████████
██████████████████████████████████████████████████████████████████████
`

// Print the large logo of this cli tool.
// The logo consists of the Sakamichi-mark and this cli name(SGI).
// The format is defined at [CLI_LOGO].
//
// This function uses dimiron1/banner package.
// To get detailed information, please see: https://github.com/dimiro1/banner
func PrintLogo() {
	banner.Init(colorable.NewColorableStdout(), true, true, bytes.NewBufferString(CLI_LOGO))
}
