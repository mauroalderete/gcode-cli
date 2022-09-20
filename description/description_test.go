package description

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {

	targets := map[string]struct {
		source        io.Reader
		filename      string
		isSuccessfull bool
	}{
		"file without content":  {strings.NewReader(""), "somefile.gcode", false},
		"stdin without content": {strings.NewReader(""), "", false},
		"file with content":     {strings.NewReader("some content"), "somefile.gcode", true},
		"stdin with content":    {strings.NewReader("some content"), "", true},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {
			desc, err := New(func(c Configurer) error {

				err := c.SetFilename(v.filename)
				if err != nil {
					return err
				}
				err = c.SetSource(v.source)
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				t.Errorf("want error nil, got %v", err)
				return
			}

			// Executes the parsing
			err = desc.Parse()
			if (err == nil) != v.isSuccessfull {
				t.Errorf("want error == %v, got error == %v: %v", v.isSuccessfull, !v.isSuccessfull, err)
				return
			}
		})
	}
}

func TestNew(t *testing.T) {

	targets := map[string]struct {
		source        io.Reader
		filename      string
		isSuccessfull bool
	}{
		"file with content nil": {nil, "somefile.gcode", false},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {
			_, err := New(func(c Configurer) error {

				c.SetFilename(v.filename)
				c.SetSource(v.source)

				return nil
			})
			if (err == nil) != v.isSuccessfull {
				t.Errorf("want error == %v, got error == %v: %v", v.isSuccessfull, !v.isSuccessfull, err)
				return
			}
		})
	}
}

func TestNewWithConfigFail(t *testing.T) {

	targets := map[string]struct {
		source        io.Reader
		filename      string
		isSuccessfull bool
	}{
		"file with content": {strings.NewReader("some content"), "somefile.gcode", false},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {
			_, err := New(func(c Configurer) error {

				c.SetSource(v.source)

				return fmt.Errorf("some error")
			})
			if (err == nil) != v.isSuccessfull {
				t.Errorf("want error == %v, got error == %v: %v", v.isSuccessfull, !v.isSuccessfull, err)
				return
			}
		})
	}
}

func TestFormatJson(t *testing.T) {
	targets := map[string]struct {
		source   string
		expected string
		failed   bool
	}{
		"invalid": {"some content", `{"filename":"","linesCount":1,"blocksCount":0,"coverage":0}`, false},
		"valid": {`/ this is a comment
		N1 G0
		N2 G1 X2
		// another comment`, `{"filename":"","linesCount":4,"blocksCount":2,"coverage":50}`, false},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {
			desc, err := New(func(c Configurer) error {

				return c.SetSource(strings.NewReader(v.source))
			})
			if err != nil {
				t.Errorf("want instance with error nil, got %v", err)
				return
			}

			err = desc.Parse()
			if err != nil {
				t.Errorf("want parse with error nil, got %v", err)
				return
			}

			result, err := desc.FormatJSON()
			if (err == nil) == v.failed {
				t.Errorf("want error == %v, got error == %v: %v", v.failed, !v.failed, err)
				return
			}

			if result != v.expected {
				t.Errorf("want %v, got %v", v.expected, result)
				return
			}
		})
	}

}

func TestFormatYaml(t *testing.T) {
	targets := map[string]struct {
		source   string
		expected string
		failed   bool
	}{
		"invalid": {"some content", `blocksCount: 0
coverage: 0
filename: ""
linesCount: 1
`, false},
		"valid": {`/ this is a comment
		N1 G0
		N2 G1 X2
		// another comment`, `blocksCount: 2
coverage: 50
filename: ""
linesCount: 4
`, false},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {
			desc, err := New(func(c Configurer) error {

				return c.SetSource(strings.NewReader(v.source))
			})
			if err != nil {
				t.Errorf("want instance with error nil, got %v", err)
				return
			}

			err = desc.Parse()
			if err != nil {
				t.Errorf("want parse with error nil, got %v", err)
				return
			}

			result, err := desc.FormatYAML()
			if (err == nil) == v.failed {
				t.Errorf("want error == %v, got error == %v: %v", v.failed, !v.failed, err)
				return
			}

			if result != v.expected {
				t.Errorf("want (%d)[%v], got (%d)[%v]", len(v.expected), v.expected, len(result), result)
				return
			}
		})
	}

}

func TestFormatTemplate(t *testing.T) {
	targets := map[string]struct {
		source   string
		template string
		expected string
		failed   bool
	}{
		"template bad input0": {"some content", "{.Filenames}{{> 3}}}", "", true},
		"template bad input1": {"some content", "{{.Filenames}}", "", true},
		"invalid":             {"some content", "{{.Filename}}\t{{.LinesCount}}\t{{.BlocksCount}}", "\t1\t0", false},
		"valid": {`/ this is a comment
		N1 G0
		N2 G1 X2
		// another comment`, "{{.Filename}}\t{{.LinesCount}}\t{{.BlocksCount}}", "\t4\t2", false},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {
			desc, err := New(func(c Configurer) error {

				return c.SetSource(strings.NewReader(v.source))
			})
			if err != nil {
				t.Errorf("want instance with error nil, got %v", err)
				return
			}

			err = desc.Parse()
			if err != nil {
				t.Errorf("want parse with error nil, got %v", err)
				return
			}

			result, err := desc.FormatTemplate(v.template)
			if (err == nil) == v.failed {
				t.Errorf("want error == %v, got error == %v: %v", v.failed, !v.failed, err)
				return
			}

			if result != v.expected {
				t.Errorf("want %v, got %v", v.expected, result)
				return
			}
		})
	}

}
