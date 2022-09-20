package description

import (
	"encoding/json"
	"fmt"
)

// FormatJSON implements [Descriptionable.FormatJSON]
func (d *Description) FormatJSON() (string, error) {

	dm := newDescriptionMarshable(*d)

	parsed, err := json.Marshal(dm)
	if err != nil {
		return "", fmt.Errorf("failed to marshall json format: %v", err)
	}

	return string(parsed), nil
}
