package describe

import (
	"io"
	"strings"
	"testing"

	"github.com/mauroalderete/gcode-cli/description"
)

func TestPrintJson(t *testing.T) {

	desc, err := description.New(func(c description.Configurer) error {
		return c.SetSource(strings.NewReader(`some content`))
	})
	if err != nil {
		t.Errorf("want instance without error, got %v", err)
		return
	}

	err = printJson(desc)
	if err != nil {
		t.Errorf("want print without error, got %v", err)
		return
	}
}

func TestPrintYaml(t *testing.T) {

	desc, err := description.New(func(c description.Configurer) error {
		return c.SetSource(strings.NewReader(`some content`))
	})
	if err != nil {
		t.Errorf("want instance without error, got %v", err)
		return
	}

	err = printYaml(desc)
	if err != nil {
		t.Errorf("want print without error, got %v", err)
		return
	}
}

func TestPrintTemplate(t *testing.T) {

	desc, err := description.New(func(c description.Configurer) error {
		return c.SetSource(strings.NewReader(`some content`))
	})
	if err != nil {
		t.Errorf("want instance without error, got %v", err)
		return
	}

	const template = ""

	err = printTemplate(desc, template)
	if err != nil {
		t.Errorf("want print without error, got %v", err)
		return
	}
}

func TestGetInput(t *testing.T) {
	targets := map[string]struct {
		stream      io.Reader
		args        []string
		fail        bool
		isSameInput bool
	}{
		"a": {nil, []string{}, true, false},
		"b": {nil, []string{"somefilethatnotexists.gcode"}, true, false},
		"c": {nil, []string{"describe_test.go"}, false, false},
		"d": {strings.NewReader(""), []string{}, false, true},
		"e": {strings.NewReader(""), []string{"somefilethatnotexists.gcode"}, true, false},
		"f": {strings.NewReader(""), []string{"describe_test.go"}, false, false},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {
			input, err := getInput(v.stream, v.args)
			if err != nil {
				if !v.fail {
					t.Errorf("want error nil, got %v", err)
					return
				}
				return
			}

			if input == v.stream && !v.isSameInput {
				t.Errorf("want input was different to stream, got same to both")
				return
			}

			if input != v.stream && v.isSameInput {
				t.Errorf("want input was equal to stream, got readers different")
				return
			}
		})
	}
}

func TestNewDescribeCommand(t *testing.T) {

	targets := map[string]struct {
		source io.Reader
		flags  map[string]string
		args   []string
		fail   bool
	}{
		"gcode describe":                            {strings.NewReader("some content"), nil, []string{}, false},
		"gcode describe --json":                     {strings.NewReader("some content"), map[string]string{"json": "true"}, []string{}, false},
		"gcode describe --yaml":                     {strings.NewReader("some content"), map[string]string{"yaml": "true"}, []string{}, false},
		"gcode describe --format":                   {strings.NewReader("some content"), map[string]string{"format": "true"}, []string{}, false},
		"gcode describe describe_test.go":           {nil, nil, []string{"describe_test.go"}, false},
		"gcode describe --json describe_test.go":    {nil, map[string]string{"json": "true"}, []string{"describe_test.go"}, false},
		"gcode describe --yaml describe_test.go":    {nil, map[string]string{"yaml": "true"}, []string{"describe_test.go"}, false},
		"gcode describe --format describe_test.go":  {nil, map[string]string{"format": "true"}, []string{"describe_test.go"}, false},
		"fgcode describe":                           {nil, nil, []string{}, true},
		"fgcode describe --json":                    {nil, map[string]string{"json": "true"}, []string{}, true},
		"fgcode describe --yaml":                    {nil, map[string]string{"yaml": "true"}, []string{}, true},
		"fgcode describe --format":                  {nil, map[string]string{"format": "true"}, []string{}, true},
		"fgcode describe describe_test.go":          {nil, nil, []string{"noexists.gcode"}, true},
		"fgcode describe --json describe_test.go":   {nil, map[string]string{"json": "true"}, []string{"noexists.gcode"}, true},
		"fgcode describe --yaml describe_test.go":   {nil, map[string]string{"yaml": "true"}, []string{"noexists.gcode"}, true},
		"fgcode describe --format describe_test.go": {nil, map[string]string{"format": "true"}, []string{"noexists.gcode"}, true},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {

			cmd := NewDescribeCommand()
			for fk, fv := range v.flags {
				cmd.Flags().Set(fk, fv)
			}
			cmd.SetIn(v.source)

			err := cmd.RunE(cmd, v.args)
			if (err != nil) != v.fail {
				t.Errorf("want error %v, got error %v: %v", v.fail, !v.fail, err)
				return
			}
		})
	}
}
