// Package version is the subcommand to show the version number of gcode-cli.
package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewVersionCommand creates an instance of *cobra.Command with behaviour to print the version number.
//
// The version number displayed is the recovery from the number version set in the root command,
//
// For more information on how to configure with a specific version value, please refer to the build options documentation.
func NewVersionCommand() *cobra.Command {

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of gcode-cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "%s version %s\n", cmd.Root().Name(), cmd.Root().Version)
		},
	}

	return versionCmd
}
