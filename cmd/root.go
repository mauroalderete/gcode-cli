// Package cmd contains the handler to configure and execute the subcommands and flags.
package cmd

import (
	"github.com/mauroalderete/gcode-cli/cmd/describe"
	"github.com/mauroalderete/gcode-cli/cmd/skew"
	cmdVersion "github.com/mauroalderete/gcode-cli/cmd/version"
	"github.com/spf13/cobra"
)

type teminator func(code int)

// NewRootCommand make a root command instance
//
// Recives a version number
func NewRootCommand(version string) *cobra.Command {

	// Configure the root command
	var cmd = &cobra.Command{
		Use:   "gcode-cli",
		Short: "Gcode command-line tool to oparate with gcode files",
		Long: `gcode-cli is a command-line tool that helps you to apply massive operations in your gcode files.
You can make skew corrections, translations or checksum integrity.`,
		Version: version,
	}

	// init Initializes and configures command
	cmd.AddCommand(cmdVersion.NewVersionCommand())
	cmd.AddCommand(describe.NewDescribeCommand())
	cmd.AddCommand(skew.NewSkewCommand())

	return cmd
}

// Execute executes the root command passed by the parameter
func Execute(cmd *cobra.Command, exit teminator) {
	err := cmd.Execute()
	if err != nil {
		exit(1)
	}
}
