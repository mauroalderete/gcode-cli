package skewer

import (
	"github.com/mauroalderete/gcode-core/gcode"
)

type Skewer struct {
	source []gcode.Gcoder
}

func (s *Skewer) Source() []gcode.Gcoder {
	return s.source
}

func (s *Skewer) SkewXY(ratio float32) error {
	return nil
}

func (s *Skewer) SkewXZ(ratio float32) error {
	return nil
}

func (s *Skewer) SkewYZ(ratio float32) error {
	return nil
}
