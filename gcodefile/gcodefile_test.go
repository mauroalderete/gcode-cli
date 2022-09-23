package gcodefile

import (
	"io"
	"strings"
	"testing"
)

func TestNewFromReader(t *testing.T) {

	targets := map[string]struct {
		source io.Reader
		expect bool
	}{
		"a": {strings.NewReader(""), true},
		"b": {strings.NewReader(";aaaaa"), true},
		"c": {strings.NewReader("G1 X2"), true},
		"d": {strings.NewReader("G1 X2\n;aaaa"), true},
		"e": {strings.NewReader("G1 X2\nG1X2"), false},
		"f": {strings.NewReader("G1 X2\n\n\nG1 X2"), true},
		"g": {strings.NewReader("G1 X2\nG1X2"), false},
		"h": {strings.NewReader("1 X2\nG1X2"), false},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {
			gf, err := NewFromReader(v.source)
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

func TestGcodeFile_Refresh(t *testing.T) {

}

func TestGcodeFile_Update(t *testing.T) {

}

func TestGcodeFile_SaveFile(t *testing.T) {

}
