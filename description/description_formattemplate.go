package description

import (
	"bytes"
	"fmt"
	"io"
	templatePkg "text/template"
)

// FormatTemplate implements [Descriptionable.FormatTemplate]
func (d *Description) FormatTemplate(template string) (string, error) {

	tpl, err := templatePkg.New("description").Parse(template)
	if err != nil {
		return "", fmt.Errorf("failed prepare template parsing: %v", err)
	}

	dm := newDescriptionMarshable(*d)

	buf := new(bytes.Buffer)

	err = tpl.Execute(buf, dm)
	if err != nil {
		return "", fmt.Errorf("failed to parse the template: %v", err)
	}

	data, err := io.ReadAll(buf)
	if err != nil {
		return "", fmt.Errorf("failed to get the output: %v", err)
	}

	return string(data), nil
}
