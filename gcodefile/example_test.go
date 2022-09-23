package gcodefile_test

import (
	"fmt"
	"strings"

	"github.com/mauroalderete/gcode-cli/gcodefile"
)

func ExampleNewFromReader() {

	source := strings.NewReader(";some gcode content")

	gf, err := gcodefile.NewFromReader(source)
	if err != nil {
		fmt.Printf("failed instance gcodefile: %v", err)
		return
	}

	if gf == nil {
		fmt.Printf("some was wrong with gcodefile instance")
		return
	}

	fmt.Printf("the gcode file is available")

	//Output: the gcode file is available
}

func ExampleNewFromFile() {

	path := "some path to gcode file"

	gf, err := gcodefile.NewFromFile(path)
	if err != nil {
		fmt.Printf("failed instance gcodefile: %v", err)
		return
	}

	if gf == nil {
		fmt.Printf("some was wrong with gcodefile instance")
		return
	}

	fmt.Printf("the gcode file is available")

	//Output: failed instance gcodefile: failed open file in 'some path to gcode file': open some path to gcode file: no such file or directory
}
