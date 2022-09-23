package gcodefile_test

import (
	"fmt"
	"io"
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

func ExampleGcodeFile_Gcodes() {

	source := strings.NewReader("G1 X0")

	gf, err := gcodefile.NewFromReader(source)
	if err != nil {
		fmt.Printf("failed instance gcodefile: %v", err)
		return
	}

	gcodes := gf.Gcodes()

	fmt.Printf("there are %d blocks parsed", len(gcodes))

	//Output: there are 1 blocks parsed
}

func ExampleGcodeFile_Source() {

	source := strings.NewReader("G1 X0")

	gf, err := gcodefile.NewFromReader(source)
	if err != nil {
		fmt.Printf("failed instance gcodefile: %v", err)
		return
	}

	content, err := io.ReadAll(gf.Source())
	if err != nil {
		fmt.Printf("failed to read the source: %v", err)
		return
	}

	fmt.Printf("the source contain %s", content)

	//Output: the source contain G1 X0
}
