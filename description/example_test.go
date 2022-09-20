package description_test

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/mauroalderete/gcode-cli/description"
)

func ExampleDescription_Filename() {

	file := "some_printable_object.gcode"

	// Creates a new instance of Description
	desc, err := description.New(func(c description.Configurer) error {

		// Configures the filename to will load
		c.SetFilename(file)

		// Configures the required source
		c.SetSource(bufio.NewReader(nil))

		return nil
	})
	if err != nil {
		fmt.Printf("failed to prepare the file: %v", err)
		return
	}

	// Checks if the filename configured match with expected
	if desc.Filename() == file {
		fmt.Printf("file loaded")
	} else {
		fmt.Printf("file missing")
	}

	// Output: file loaded
}

func ExampleDescription_LinesCount() {

	source := strings.NewReader(`// this is a comment
	N1 G0
	N2 G1 X2
	// another comment`)

	// Creates a new instance of Description
	desc, err := description.New(func(c description.Configurer) error {

		// Configures the required source
		c.SetSource(source)

		return nil
	})
	if err != nil {
		fmt.Printf("failed to prepare the file: %v", err)
		return
	}

	// Executes the parsing
	err = desc.Parse()
	if err != nil {
		fmt.Printf("some was wrong to try parse the source: %v", err)
	}

	// Shows the count of lines
	fmt.Printf("the source contains %d lines", desc.LinesCount())

	// Output: the source contains 4 lines
}

func ExampleDescription_BlocksCount() {

	source := strings.NewReader(`// this is a comment
	N1 G0
	N2 G1 X2
	// another comment`)

	// Creates a new instance of Description
	desc, err := description.New(func(c description.Configurer) error {

		// Configures the required source
		c.SetSource(source)

		return nil
	})
	if err != nil {
		fmt.Printf("failed to prepare the file: %v", err)
		return
	}

	// Executes the parsing
	err = desc.Parse()
	if err != nil {
		fmt.Printf("some was wrong to try parse the source: %v", err)
	}

	// Shows the count of blocks
	fmt.Printf("the source contains %d blocks", desc.BlocksCount())

	// Output: the source contains 2 blocks
}

func ExampleDescription_Parse() {

	source := strings.NewReader(`// this is a comment
	N1 G0
	N2 G1 X2
	// another comment`)

	// Creates a new instance of Description
	desc, err := description.New(func(c description.Configurer) error {

		// Configures the required source
		c.SetSource(source)

		return nil
	})
	if err != nil {
		fmt.Printf("failed to prepare the file: %v", err)
		return
	}

	// Executes the parsing
	err = desc.Parse()
	if err != nil {
		fmt.Printf("some was wrong to try parse the source: %v", err)
	}

	fmt.Printf("the source parsed succefull")

	// Output: the source parsed succefull
}
