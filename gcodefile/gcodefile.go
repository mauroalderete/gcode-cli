package gcodefile

import (
	"io"

	"github.com/mauroalderete/gcode-core/gcode"
)

type GcodeFile struct {
	source io.Reader
	gcodes []gcode.Gcoder
}

func (gf *GcodeFile) Gcodes() []gcode.Gcoder {
	return gf.gcodes
}

func (gf *GcodeFile) Refresh() error {
	return nil
}

func (gf *GcodeFile) Update() (io.Reader, error) {
	return nil, nil
}

func (gf *GcodeFile) SaveFile(path string) error {
	return nil
}

func NewFromFile(source string) (*GcodeFile, error) {
	return nil, nil
}

func NewFromReader(source io.Reader) (*GcodeFile, error) {
	return nil, nil
}
