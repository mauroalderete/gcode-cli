//go:build !test

package main

import (
	"os"

	"github.com/mauroalderete/gcode-cli/cmd"
)

// Stores version. It is set from -ldflags in build time
var version string

func main() {
	rootCmd := cmd.NewRootCommand(version)
	cmd.Execute(rootCmd, os.Exit)
}
