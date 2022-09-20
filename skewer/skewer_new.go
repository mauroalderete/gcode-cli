package skewer

import (
	"fmt"

	"github.com/mauroalderete/gcode-core/gcode"
)

// Configurer is a interface that contains all functions that allows configure a new [Skewer] instance.
type Configurer interface {
	// SetSource allows set the [gcode.Gcoder] elements to fix them.
	SetSource(source []gcode.Gcoder) error
}

// configure implement [Configurer] interface.
//
// Although of the public methods defined by [Configurer],
// configure can validate the options values and stores a series of clouser callbacks
// that recive a [Skewer] instance.
// This clouser callbacks operate with the [Skewer] instance and option values to apply the configuration.
type configure struct {
	pool []func(*Skewer)
}

// SetSource implements [Configurer.SetSource], return error if source is null
func (c *configure) SetSource(source []gcode.Gcoder) error {

	if source == nil {
		return fmt.Errorf("skewer requires a source: source is empty")
	}

	c.pool = append(c.pool, func(d *Skewer) {
		d.source = source
	})

	return nil
}

func New(options ...func(Configurer) error) (*Skewer, error) {

	return nil, nil
}
