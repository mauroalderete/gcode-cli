package skewer_test

import (
	"fmt"

	"github.com/mauroalderete/gcode-cli/skewer"
	"github.com/mauroalderete/gcode-core/gcode"
)

func ExampleNew() {

	// Prepares a source with gcode blocks
	source := make([]gcode.Gcoder, 0)

	// Instances a new skewer object with source loaded
	skw, err := skewer.New(func(c skewer.Configurer) error {
		return c.SetSource(source)
	})

	// Checks if was some problem while try to instance
	if err != nil {
		fmt.Printf("failed instance skewer: %v", err)
		return
	}

	// Checks if there is a skewer instance live
	if skw == nil {
		fmt.Printf("skewer is empty")
		return
	}

	fmt.Printf("skewer ready to process!")

	//Output: skewer ready to process!
}
