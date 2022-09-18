// Package describe provides a implementation that parse a gcode file to recompile staticts and metrics.
package description

import (
	"io"
)

// Descriptionable is a main interface to handle a [Description] of a gcode file
type Descriptionable interface {
	// Filename returns the filename of the gcode file input
	// or a string empty if the input from stdin.
	Filename() string

	// LinesCount return the number of lines contains in input,
	// this includes comments.
	LinesCount() int

	// BlocksCount returns the number of lines that are a gcode block valid.
	BlocksCount() int

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
	filename    string
	linesCount  int
	blocksCount int
}

// Filename implements [Descriptionable.Filename]
func (d *Description) Filename() string {
	return d.filename
}

// LinesCount implements [Descriptionable.LinesCount]
func (d *Description) LinesCount() int {
	return d.linesCount
}

// BlocksCount implements [Descriptionable.BlocksCount]
func (d *Description) BlocksCount() int {
	return d.blocksCount
}
