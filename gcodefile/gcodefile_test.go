package gcodefile

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/mauroalderete/gcode-core/block/gcodeblock"
)

func TestNewFromReader(t *testing.T) {

	t.Run("nil source", func(t *testing.T) {
		gf, err := NewFromReader(nil)
		if err == nil {
			t.Errorf("want an error, got error nil")
			return
		}

		if gf != nil {
			t.Errorf("want an gcodefile instance nil, got an gcodefile instanced: %v", gf)
			return
		}
	})

	t.Run("source with values", func(t *testing.T) {
		targets := map[string]struct {
			source string
			expect bool
		}{
			"a": {"", true},
			"b": {";aaaaa", true},
			"c": {"G1 X2", true},
			"d": {"G1 X2\n;aaaa", true},
			"e": {"G1 X2\nG1X2", false},
			"f": {"G1 X2\n\n\nG1 X2", true},
			"g": {"G1 X2\nG1X2", false},
			"h": {"1 X2\nG1X2", false},
		}

		for k, v := range targets {
			t.Run(k, func(t *testing.T) {

				gf, err := NewFromReader(strings.NewReader(v.source))
				if err != nil && v.expect {
					t.Errorf("want error nil, got error %v", err)
					return
				}
				if gf == nil && v.expect {
					t.Errorf("want gcodefile not nil, got gcodefile nil")
					return
				}
				if err == nil && !v.expect {
					t.Errorf("want error not nil, got error nil")
					return
				}
				if gf != nil && !v.expect {
					t.Errorf("want gcodefile nil by error, got gcodefile nil with err nil")
					return
				}

				if !v.expect {
					return
				}

				src, err := io.ReadAll(gf.Source())
				if err != nil {
					t.Errorf("fail to read source: want error nil, got error %v", err)
					return
				}

				if v.source != string(src) {
					t.Errorf("target source is not equal to source stored: want error nil, got error %v", err)
					return
				}
			})
		}
	})
}

func TestNewFromFile(t *testing.T) {
	targets := map[string]struct {
		path   string
		expect bool
	}{
		"valid file":   {"../data/test.gcode", true},
		"invalid file": {"somewrongfile", false},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {
			gf, err := NewFromFile(v.path)
			if err != nil && v.expect {
				t.Errorf("want error nil, got error %v", err)
				return
			}
			if gf == nil && v.expect {
				t.Errorf("want gcodefile not nil, got gcodefile nil")
				return
			}
			if err == nil && !v.expect {
				t.Errorf("want error not nil, got error nil")
				return
			}
			if gf != nil && !v.expect {
				t.Errorf("want gcodefile nil by error, got gcodefile nil with err nil")
				return
			}
		})
	}
}

func TestGcodeFile_Gcodes(t *testing.T) {

}

func TestGcodeFile_Update(t *testing.T) {

	t.Run("double refresh", func(t *testing.T) {
		targets := map[string]struct {
			source string
		}{
			"a": {"G1 X1"},
		}

		for k, v := range targets {
			t.Run(k, func(t *testing.T) {

				gf, err := NewFromReader(strings.NewReader(v.source))
				if err != nil {
					t.Errorf("want error nil, got error %v", err)
					return
				}

				if gf == nil {
					t.Errorf("want gcodefile not nil, got gcodefile nil")
					return
				}

				gc := gf.Gcodes()

				block, err := gcodeblock.Parse("G2 X2")
				if err != nil {
					t.Errorf("failed instance a mock gcodeblock: %v", err)
					return
				}

				gc = append(gc, *block)

				gf.Update(gc)

				gc2 := gf.Gcodes()

				if len(gc2) != 2 {
					t.Errorf("want 2 gcodes updated, got %d", len(gc2))
					return
				}
			})
		}
	})

}

func TestGcodeFile_SaveFile(t *testing.T) {

	t.Run("valid save", func(t *testing.T) {
		testFilePath := "../data/testsave.gcode"

		gf, err := NewFromReader(strings.NewReader("G1 X1"))
		if err != nil {
			t.Errorf("want error nil, got error %v", err)
			return
		}

		if gf == nil {
			t.Errorf("want gcodefile not nil, got gcodefile nil")
			return
		}

		err = gf.SaveFile(testFilePath)
		if err != nil {
			t.Errorf("failed save file: want error nil, got error %v", err)
			return
		}

		_, err = os.Open(testFilePath)
		if err != nil {
			t.Errorf("failed reopen saved file: want error nil, got error %v", err)
			return
		}

		err = os.Remove(testFilePath)
		if err != nil {
			t.Errorf("failed remove saved file: want error nil, got error %v", err)
			return
		}
	})

	t.Run("invalid save", func(t *testing.T) {
		testFilePath := "../data/bad folder/testsave.gcode"

		gf, err := NewFromReader(strings.NewReader("G1 X1"))
		if err != nil {
			t.Errorf("want error nil, got error %v", err)
			return
		}

		if gf == nil {
			t.Errorf("want gcodefile not nil, got gcodefile nil")
			return
		}

		err = gf.SaveFile(testFilePath)
		if err == nil {
			t.Errorf("want an error, got error nil")
			return
		}
	})

}
