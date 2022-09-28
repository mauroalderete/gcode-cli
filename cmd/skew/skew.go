// Package skew contains the subcommand to apply skew correction
package skew

import (
	"github.com/mauroalderete/gcode-cli/cmd/skew/xy"
	"github.com/spf13/cobra"
)

// NewSkewCommand creates an instance of *cobra.Command with the behaviour
func NewSkewCommand() *cobra.Command {

	// Defines the command skew
	var cmd = &cobra.Command{
		Use:   "skew",
		Short: "Fix the skew over a plane from angle since",
		Long: `Apply an angle skew correction over some plane XY, ZY, XZ from gcode [FILE]
With no FILE, or when FILE is -, read standard input.
`,
	}

	cmd.AddCommand(xy.NewXYCommand())

	return cmd
}
