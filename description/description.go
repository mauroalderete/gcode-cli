// Package describe provides a implementation that parse a gcode and recompile staticts and metrics.
package description

import (
	"io"
)

// Descriptionable is a main interface to handle a Description of a gcode file
type Descriptionable interface {

	// Parse evaluates the gcode file stored to fill the internal fields.
	Parse() error

	// FormatJSON returns the Descriptionable instance value in YAML format or an error trying to do it.
	FormatJSON() (string, error)

	// FormatYAML returns the Descriptionable instance value in YAML format or an error trying to do it.
	FormatYAML() (string, error)

	// FormatTemplate returns the Descriptionable instance value using a Go template format or an error trying to do it.
	FormatTemplate() (string, error)
}

// Description implements Descriptionable interface
type Description struct {
	source      io.Reader
	Filename    string `json:"filename"`
	LinesCount  int    `json:"linesCount"`
	BlocksCount int    `json:"blocksCount"`
}
