// Package skew contains the subcommand to apply skew correction
package xy

import (
	"fmt"
	"io"
	"os"

	"github.com/mauroalderete/gcode-cli/skewer"
	"github.com/spf13/cobra"
)

// flag[T interface{}] allows define a flag and binding its value
type flag[T interface{}] struct {
	name  string
	value T
}

// NewSkewCommand creates an instance of *cobra.Command with the behaviour
func NewXYCommand() *cobra.Command {

	// Initializes flags bindable
	options := struct {
		degree flag[float32]
		radian flag[float32]
		ratio  flag[float32]
	}{
		degree: flag[float32]{name: "degree", value: 0},
		radian: flag[float32]{name: "radian", value: 0},
		ratio:  flag[float32]{name: "ratio", value: 0},
	}

	// Defines the command describe
	var cmd = &cobra.Command{
		Use:   "xy [FILE]",
		Short: "Fix skew from angle",
		Long: `Apply an angle skew correction from gcode [FILE]
With no FILE, or when FILE is -, read standard input.
`,
		Args: cobra.MaximumNArgs(1),
	}

	// Load the falgs using options instance to binding the values
	// Load the falgs using options instance to binding the values
	cmd.Flags().Float32Var(&options.degree.value, options.degree.name, options.degree.value, "Skew value in degrees")
	cmd.Flags().Float32Var(&options.radian.value, options.radian.name, options.radian.value, "Skew value in radians")
	cmd.Flags().Float32Var(&options.ratio.value, options.ratio.name, options.ratio.value, "Skew value in ratio")

	cmd.MarkFlagsMutuallyExclusive(
		options.degree.name,
		options.radian.name,
		options.ratio.name,
	)

	// Loads the handler of command.
	// Consumes options to determine the correct behaivour
	cmd.RunE = func(cmd *cobra.Command, args []string) error {

		input, err := getInput(cmd.InOrStdin(), args)
		if err != nil {
			return err
		}

		// gcode file struct
		//gets origin of source (file or stdin) => io.Reader
		//	probably two news
		//parses io.Reader => []gcode.Gcoder
		//
		// apply changes to each gcoder that needed it
		//
		//generates new file from []gcode.Gcoder
		//generates maintaining comments??
		//
		//save new file / flag output file

		_, err = skewer.New(func(c skewer.Configurer) error {
			return c.SetSource(nil)
		})
		if err != nil {
			return fmt.Errorf("failed to instanciate Skewer: %v", err)
		}

		return nil
	}

	return cmd
}

// getInput return the correct input of gcode file as io.Reader instance according to arguments available
func getInput(stdio io.Reader, args []string) (io.Reader, error) {

	var input io.Reader = stdio

	if len(args) > 0 && args[0] != "-" {
		file, err := os.Open(args[0])
		if err != nil {
			return nil, fmt.Errorf("failed open file: %v", err)
		}
		input = file
	}

	if input == nil {
		return nil, fmt.Errorf("input missing")
	}

	return input, nil
}
