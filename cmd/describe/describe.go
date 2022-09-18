// Package describe is the subcommand to show the metadata from gcode file.
package describe

import (
	"fmt"
	"io"
	"os"

	descriptionModule "github.com/mauroalderete/gcode-cli/description"
	"github.com/spf13/cobra"
)

// flag[T interface{}] allows define a flag and binding its value
type flag[T interface{}] struct {
	name  string
	value T
}

// NewDescribeCommand creates an instance of *cobra.Command with the behaviour
// to print metadata from gcode file passed as argument or stdin.
func NewDescribeCommand() *cobra.Command {

	// Initializes flags bindable
	options := struct {
		json     flag[bool]
		yaml     flag[bool]
		template flag[string]
	}{
		json:     flag[bool]{name: "json", value: false},
		yaml:     flag[bool]{name: "yaml", value: false},
		template: flag[string]{name: "format", value: ""},
	}

	// Defines the command describe
	var cmd = &cobra.Command{
		Use:   "describe [FILE]",
		Short: "Print metadata from gcode [FILE]",
		Long:  `With no FILE, or when FILE is -, read standar input.`,
		Args:  cobra.MaximumNArgs(1),
	}

	// Load the falgs using options instance to binding the values
	cmd.Flags().BoolVar(&options.json.value, options.json.name, options.json.value, "Output in JSON")
	cmd.Flags().BoolVar(&options.yaml.value, options.yaml.name, options.yaml.value, "Output in YAML")
	cmd.Flags().StringVar(&options.template.value, options.template.name, options.template.value, "Pretty-print containers using a Go template")
	cmd.MarkFlagsMutuallyExclusive(
		options.json.name,
		options.yaml.name,
		options.template.name,
	)

	// Loads the handler of command.
	// Consumes options to determine the correct behaivour
	cmd.RunE = func(cmd *cobra.Command, args []string) error {

		filename, input, err := getInput(cmd.InOrStdin(), args)
		if err != nil {
			return err
		}

		description, err := descriptionModule.New(func(cn descriptionModule.Configurer) error {
			cn.SetFilename(filename)
			cn.SetSource(input)
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to instanciate Description from input: %v", err)
		}

		err = description.Parse()
		if err != nil {
			return fmt.Errorf("failed to try parse input: %v", err)
		}

		// Parses flags to determine the format to print the Description instance.
		if options.json.value {
			return printJson(description)
		}

		if options.yaml.value {
			return printYaml(description)
		}

		return printTemplate(description, options.template.value)
	}

	return cmd
}

// getInput return the correct input of gcode file as io.Reader instance according to arguments available
func getInput(stdio io.Reader, args []string) (string, io.Reader, error) {

	var input io.Reader = stdio

	if len(args) > 0 && args[0] != "-" {
		file, err := os.Open(args[0])
		if err != nil {
			return "", nil, fmt.Errorf("failed open file: %v", err)
		}
		input = file
	}

	if input == nil {
		return "", nil, fmt.Errorf("input missing")
	}

	return args[0], input, nil
}

// printJson prints on the stdout the Description instance in JSON format
func printJson(d descriptionModule.Descriptionable) error {

	parsed, err := d.FormatJSON()
	if err != nil {
		return fmt.Errorf("failed to get description in json format: %v", err)
	}

	fmt.Printf("%s\n", parsed)
	return nil
}

// printYaml prints on the stdout the Description instance in YAML format
func printYaml(d descriptionModule.Descriptionable) error {
	return nil
}

// printTempalte prints on the stdout the Description instance using a Go template format
func printTemplate(d descriptionModule.Descriptionable, template string) error {
	return nil
}