package description

import (
	"fmt"
	"io"
)

// ConfigurerNew is a interface that contains all functions that allows configure a new Description instance.
type ConfigurerNew interface {
	// SetSource allows set the source content to get the metadatas.
	SetSource(source io.Reader) error
}

// configurationPoolNew implement ConfigurerNew interface.
//
// Although of the public methods defined by ConfigurerNew,
// configurationPoolNew can validate the options values and stores a series of clouser callbacks
// that recive a Describe instance.
// This clouser callbacks operate with the Desribe instance and option values to apply the configuration.
type configurationPoolNew struct {
	pool []func(*Description) error
}

// SetSource implements ConfigurerNew.SetSource, return error if source is null
func (c *configurationPoolNew) SetSource(source io.Reader) error {

	if source == nil {
		return fmt.Errorf("describe requires a source: source is empty")
	}

	c.pool = append(c.pool, func(d *Description) error {
		d.source = source
		return nil
	})

	return nil
}

// New creates a instance of Description with all statics and metrics parsed from a gcode file.
//
// It recives a list of configurations callbacks to set a io.Reader instance as the source of gcode file.
//
// An error is returned if something was wrong while try initialize the Description instance.
func New(options ...func(ConfigurerNew) error) (*Description, error) {

	describe := &Description{}

	configurer := &configurationPoolNew{}

	for _, opt := range options {
		err := opt(configurer)
		if err != nil {
			return nil, fmt.Errorf("failed to load Describe option: %v", err)
		}
	}

	for _, config := range configurer.pool {
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
