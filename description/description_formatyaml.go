package description

import (
	"fmt"

	"github.com/ghodss/yaml"
)

type descriptionYaml struct {
	Filename    string `yaml:"filename"`
	LinesCount  int    `yaml:"linesCount"`
	BlocksCount int    `yaml:"blocksCount"`
}

// FormatYAML implements [Descriptionable.FormatYAML]
func (d *Description) FormatYAML() (string, error) {

	descriptionMarshable := &descriptionYaml{
		Filename:    d.Filename(),
		LinesCount:  d.LinesCount(),
		BlocksCount: d.BlocksCount(),
	}

	data, err := yaml.Marshal(descriptionMarshable)
	if err != nil {
		return "", fmt.Errorf("failed to marshall json format: %v", err)
	}

	return string(data), nil
}
