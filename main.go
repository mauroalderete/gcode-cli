//go:build !test

package main

import (
	"os"

	_ "embed"

	"github.com/mauroalderete/gcode-cli/cmd"
)

//go:generate build/get_version.sh
//go:embed build/version.txt
var version string

func main() {
	rootCmd := cmd.NewRootCommand(version)
	cmd.Execute(rootCmd, os.Exit)
}
