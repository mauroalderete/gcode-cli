package description

import (
	"fmt"

	"github.com/ghodss/yaml"
)

// FormatYAML implements [Descriptionable.FormatYAML]
func (d *Description) FormatYAML() (string, error) {

	dm := newDescriptionMarshable(*d)

	data, err := yaml.Marshal(dm)
	if err != nil {
		return "", fmt.Errorf("failed to marshall json format: %v", err)
	}

	return string(data), nil
}
