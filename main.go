//go:build !test

package main

import (
	"os"

	"github.com/mauroalderete/gcode-cli/cmd"
)

var version string

func main() {
	rootCmd := cmd.NewRootCommand(version)
	cmd.Execute(rootCmd, os.Exit)
}
