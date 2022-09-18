package description

import (
	"encoding/json"
	"fmt"
)

type descriptionJson struct {
	Filename    string `json:"filename"`
	LinesCount  int    `json:"linesCount"`
	BlocksCount int    `json:"blocksCount"`
}

// FormatJSON implements [Descriptionable.FormatJSON]
func (d *Description) FormatJSON() (string, error) {

	dd := &descriptionJson{
		Filename:    d.Filename(),
		LinesCount:  d.LinesCount(),
		BlocksCount: d.BlocksCount(),
	}

	parsed, err := json.Marshal(dd)
	if err != nil {
		return "", fmt.Errorf("failed to parse description in json format: %v", err)
	}

	return string(parsed), nil
}
