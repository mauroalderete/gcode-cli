package skewer

import (
	"testing"

	"github.com/mauroalderete/gcode-core/block/gcodeblock"
)

func TestSkewXY(t *testing.T) {

	targets := map[string]struct {
		blocks []string
		ratio  float32
		expect []string
	}{
		"a": {[]string{"G1 X2 Y1"}, 0.5, []string{"G1 X1.5 Y1"}},
		"b": {[]string{"G1 X1 Y1"}, 0.5, []string{"G1 X0.5 Y1"}},
		"c": {[]string{"G1 X1.2 Y1"}, 0.5, []string{"G1 X0.7 Y1"}},
		"d": {[]string{"G1 X0 Y1"}, 0.5, []string{"G1 X-0.5 Y1"}},
		"e": {[]string{"G1 X-1 Y1"}, 0.5, []string{"G1 X-1.5 Y1"}},
		"f": {[]string{"G1 X-1.2 Y1"}, 0.5, []string{"G1 X-1.7 Y1"}},
		"g": {[]string{
			"G1 X1 Y1",
			"G1 X2 Y1",
		}, 0.5, []string{
			"G1 X0.5 Y1",
			"G1 X1.5 Y1",
		}},
	}

	for k, targetCase := range targets {
		t.Run(k, func(t *testing.T) {
			blocks := make([]gcodeblock.GcodeBlock, 0)

			for _, blockCase := range targetCase.blocks {

				b, err := gcodeblock.Parse(blockCase)
				if err != nil {
					continue
				}

				blocks = append(blocks, *b)
			}

			result, err := SkewXY(targetCase.ratio, blocks)
			if err != nil {
				t.Errorf("want not error, got error: %v", err)
				return
			}

			if len(result) != len(targetCase.expect) {
				t.Errorf("want result and target contains are same len, got result:%d expected:%d", len(result), len(targetCase.expect))
				return
			}

			for i, r := range result {
				if r.ToLine("%c %p") != targetCase.expect[i] {
					t.Errorf("expected '%s', got result '%s'", targetCase.expect[i], r.ToLine("%c %p"))
					return
				}
			}
		})
	}
}

func TestGetParameterFloat32(t *testing.T) {
	targets := map[string]struct {
		block    string
		param    byte
		expected float32
		fail     bool
	}{
		"a": {"G1 X1.2 Y3.4", 'X', 1.2, false},
		"b": {"G1 X1.2 Y3.4", 'Y', 3.4, false},
		"c": {"G1 X1.2 Y3.4", 'Z', 0, true},
		"d": {"G1 X1 Y3.4", 'X', 1, false},
		"e": {"G1 X-1 Y3.4", 'X', -1, false},
		"f": {"G1 X-1.0 Y3.4", 'X', -1, false},
	}

	for k, target := range targets {
		t.Run(k, func(t *testing.T) {
			block, err := gcodeblock.Parse(target.block)
			if err != nil {
				t.Errorf("failed to prepare the test case")
				return
			}

			param, err := getParameterFloat32(target.param, block.Parameters())
			if (err != nil) != target.fail {
				t.Errorf("expected error %v, got error %v: %v", target.fail, (err != nil), err)
				return
			}
			if err != nil {
				return
			}

			if param.Address() != target.expected {
				t.Errorf("expected value %f, got %f", target.expected, param.Address())
			}
		})
	}
}
