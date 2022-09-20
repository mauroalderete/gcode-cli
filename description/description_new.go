package description

import (
	"fmt"
	"io"
)

// Configurer is a interface that contains all functions that allows configure a new [Description] instance.
type Configurer interface {
	// SetSource allows set the source content to get the metadatas.
	SetSource(source io.Reader) error
	// SetFilename allows set the file path to the source content.
	SetFilename(filename string) error
}

// configure implement [Configurer] interface.
//
// Although of the public methods defined by [Configurer],
// configure can validate the options values and stores a series of clouser callbacks
// that recive a Describe instance.
// This clouser callbacks operate with the Desribe instance and option values to apply the configuration.
type configure struct {
	pool []func(*Description) error
}

// SetSource implements [Configurer.SetSource], return error if source is null
func (c *configure) SetSource(source io.Reader) error {

	if source == nil {
		return fmt.Errorf("describe requires a source: source is empty")
	}

	c.pool = append(c.pool, func(d *Description) error {
		d.source = source
		return nil
	})

	return nil
}

// SetFilename implements [Configurer.SetSource], return error if source is null
func (c *configure) SetFilename(filename string) error {

	c.pool = append(c.pool, func(d *Description) error {
		d.filename = filename
		return nil
	})

	return nil
}

// New creates a instance of [Description] with all statics and metrics parsed from a gcode file.
//
// It recives a list of configurations callbacks to set a io.Reader instance as the source of gcode file.
//
// An error is returned if something was wrong while try initialize the [Description] instance.
func New(options ...func(Configurer) error) (*Description, error) {

	describe := &Description{}

	cfg := &configure{}

	for _, opt := range options {
		err := opt(cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to load Describe option: %v", err)
		}
	}

	for _, config := range cfg.pool {
		err := config(describe)
		if err != nil {
			return nil, fmt.Errorf("failed to apply Describe option: %v", err)
		}
	}

	if describe.source == nil {
		return nil, fmt.Errorf("failed to instance Describe: source is required")
	}

	return describe, nil
}
