package cmd_test

import (
	"os"

	"github.com/mauroalderete/gcode-cli/cmd"
)

func Example() {
	// Makes a root command instance with version 1.0.0
	rootCmd := cmd.NewRootCommand("1.0.0")

	// Runs root command with the terminator function to handle exit process
	cmd.Execute(rootCmd, os.Exit)
}
