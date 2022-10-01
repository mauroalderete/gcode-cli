// Package skew contains the subcommand to apply skew correction
package xy

import (
	"fmt"
	"io"
	"math"
	"os"

	"github.com/mauroalderete/gcode-cli/gcodefile"
	"github.com/mauroalderete/gcode-cli/skewer"
	"github.com/spf13/cobra"
)

// flag[T interface{}] allows define a flag and binding its value
type flag[T interface{}] struct {
	name  string
	value T
}

type skewXYFlags struct {
	degree flag[float32]
	radian flag[float32]
	ratio  flag[float32]
	output flag[string]
}

// NewSkewCommand creates an instance of *cobra.Command with the behaviour
func NewXYCommand() *cobra.Command {

	// Initializes flags bindable
	flags := skewXYFlags{
		degree: flag[float32]{name: "degree", value: 0},
		radian: flag[float32]{name: "radian", value: 0},
		ratio:  flag[float32]{name: "ratio", value: 0},
		output: flag[string]{name: "output", value: ""},
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
	cmd.Flags().Float32Var(&flags.degree.value, flags.degree.name, flags.degree.value, "Skew value in degrees")
	cmd.Flags().Float32Var(&flags.radian.value, flags.radian.name, flags.radian.value, "Skew value in radians")
	cmd.Flags().Float32Var(&flags.ratio.value, flags.ratio.name, flags.ratio.value, "Skew value in ratio")
	cmd.Flags().StringVar(&flags.output.value, flags.output.name, flags.output.value, "File output to store the gcode fixed")

	cmd.MarkFlagsMutuallyExclusive(
		flags.degree.name,
		flags.radian.name,
		flags.ratio.name,
	)

	cmd.RunE = func(cmd *cobra.Command, args []string) error {

		ratio, err := getRatio(cmd, flags)
		if err != nil {
			return err
		}

		input, err := getInput(cmd.InOrStdin(), args)
		if err != nil {
			return err
		}

		gfile, err := gcodefile.NewFromReader(input)
		if err != nil {
			return fmt.Errorf("failed instance a gcode file: %v", err)
		}

		blocks := gfile.Gcodes()

		blocksSkewed, err := skewer.SkewXY(ratio, blocks)
		if err != nil {
			return fmt.Errorf("failed apply skew transformation: %v", err)
		}

		gfile.Update(blocksSkewed)

		if flags.output.value == "" {
			toPrint, err := io.ReadAll(gfile.Source())
			if err != nil {
				return fmt.Errorf("failed print gcode source transformated: %v", err)
			}
			cmd.OutOrStdout().Write(toPrint)
			return nil
		}

		err = gfile.SaveFile(flags.output.value)
		if err != nil {
			return fmt.Errorf("failed save gcode source transformated: %v", err)
		}

		return nil
	}

	return cmd
}

func getRatio(cmd *cobra.Command, f skewXYFlags) (float32, error) {

	if cmd.Flags().Lookup(f.ratio.name).Changed {
		return f.ratio.value, nil
	}

	if cmd.Flags().Lookup(f.degree.name).Changed {
		rad := deg2rad(f.degree.value)
		return rad2rat(rad), nil
	}

	if cmd.Flags().Lookup(f.radian.name).Changed {
		return rad2rat(f.radian.value), nil
	}

	return 0, fmt.Errorf("must indicate a measure for skew correction: use --ratio --degree --radian flags")
}

func deg2rad(deg float32) float32 {
	for deg > 360 {
		deg -= 360
	}
	for deg < 0 {
		deg += 360
	}

	rad := deg * math.Pi / 180

	return rad
}

func rad2rat(rad float32) float32 {
	rat := float32(math.Tan(float64(rad)))

	return rat
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

	return input, nil
}
