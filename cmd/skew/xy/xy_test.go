package xy

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestDeg2rad(t *testing.T) {
	targets := map[string]struct {
		degree float32
		radian float32
	}{
		"left":          {180, math.Pi},
		"right":         {0, 0},
		"up_positive":   {90, float32(1) / float32(2) * math.Pi},
		"up_positive2":  {90 + 360, float32(1) / float32(2) * math.Pi},
		"down_negative": {-90, float32(3) / float32(2) * math.Pi},
		"down_positive": {90 + 180, float32(3) / float32(2) * math.Pi},
	}

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {

			rad := deg2rad(v.degree)

			if rad != v.radian {
				t.Errorf("expect %v, got %v", v.radian, rad)
				return
			}

		})
	}
}

func TestRad2Rat(t *testing.T) {
	targets := map[string]struct {
		radian float32
		ratio  float32
	}{
		"a": {deg2rad(0), 0},
		"b": {deg2rad(10), 0.17632699},
		"c": {deg2rad(20), 0.36397025},
		"d": {deg2rad(30), 0.57735026},
		"e": {deg2rad(40), 0.8390997},
		"f": {deg2rad(45), 1},
		"g": {deg2rad(50), 1.1917536},
		"h": {deg2rad(60), 1.7320509},
		"i": {deg2rad(135), -1},
		"j": {deg2rad(180), 0.000000087},
	}

	epsilon := float32(0.0000001)

	for k, v := range targets {
		t.Run(k, func(t *testing.T) {
			ratio := rad2rat(v.radian)

			if ratio-v.ratio > epsilon {
				t.Errorf("expect epsilon %v or less, got %v diff", epsilon, ratio-v.ratio)
			}
		})
	}
}

func TestGetRatio(t *testing.T) {

	// Initializes flags bindable
	flags := skewXYFlags{
		degree: flag[float32]{name: "degree", value: 0},
		radian: flag[float32]{name: "radian", value: 0},
		ratio:  flag[float32]{name: "ratio", value: 0},
	}

	t.Run("flags missing", func(t *testing.T) {

		cmd := &cobra.Command{}
		cmd.Flags().Float32Var(&flags.degree.value, flags.degree.name, flags.degree.value, "Skew value in degrees")
		cmd.Flags().Float32Var(&flags.radian.value, flags.radian.name, flags.radian.value, "Skew value in radians")
		cmd.Flags().Float32Var(&flags.ratio.value, flags.ratio.name, flags.ratio.value, "Skew value in ratio")

		_, err := getRatio(cmd, flags)
		if err == nil {
			t.Errorf("expected an error, got error nil")
		}
	})

	t.Run("ratio flag", func(t *testing.T) {

		cmd := &cobra.Command{}
		cmd.Flags().Float32Var(&flags.degree.value, flags.degree.name, flags.degree.value, "Skew value in degrees")
		cmd.Flags().Float32Var(&flags.radian.value, flags.radian.name, flags.radian.value, "Skew value in radians")
		cmd.Flags().Float32Var(&flags.ratio.value, flags.ratio.name, flags.ratio.value, "Skew value in ratio")

		cmd.Flags().Lookup(flags.ratio.name).Value.Set("1")
		cmd.Flags().Lookup(flags.ratio.name).Changed = true

		ratio, err := getRatio(cmd, flags)
		if err != nil {
			t.Errorf("expected error nil, got error %v", err)
		}
		epsilon := float32(0.00001)
		if ratio-1 > epsilon {
			t.Errorf("expected ratio 1, got ratio %f", ratio)
		}
	})

	t.Run("radian flag", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().Float32Var(&flags.degree.value, flags.degree.name, flags.degree.value, "Skew value in degrees")
		cmd.Flags().Float32Var(&flags.radian.value, flags.radian.name, flags.radian.value, "Skew value in radians")
		cmd.Flags().Float32Var(&flags.ratio.value, flags.ratio.name, flags.ratio.value, "Skew value in ratio")

		cmd.Flags().Lookup(flags.radian.name).Value.Set(fmt.Sprintf("%f", deg2rad(10)))
		cmd.Flags().Lookup(flags.radian.name).Changed = true

		ratio, err := getRatio(cmd, flags)
		if err != nil {
			t.Errorf("expected error nil, got error %v", err)
		}
		epsilon := float32(0.00001)
		expect := float32(0.17632699)
		if ratio-expect > epsilon {
			t.Errorf("expected ratio %f, got ratio %f", expect, ratio)
		}
	})

	t.Run("degree flag", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().Float32Var(&flags.degree.value, flags.degree.name, flags.degree.value, "Skew value in degrees")
		cmd.Flags().Float32Var(&flags.radian.value, flags.radian.name, flags.radian.value, "Skew value in radians")
		cmd.Flags().Float32Var(&flags.ratio.value, flags.ratio.name, flags.ratio.value, "Skew value in ratio")

		cmd.Flags().Lookup(flags.degree.name).Value.Set("45")
		cmd.Flags().Lookup(flags.degree.name).Changed = true

		ratio, err := getRatio(cmd, flags)
		if err != nil {
			t.Errorf("expected error nil, got error %v", err)
		}
		epsilon := float32(0.00001)
		if ratio-1 > epsilon {
			t.Errorf("expected ratio 1, got ratio %f", ratio)
		}
	})

}

type stdoutMock struct {
	content string
}

func (s *stdoutMock) Write(p []byte) (n int, err error) {
	s.content = string(p)
	return len(p), err
}

func TestNewXYCommand(t *testing.T) {

	t.Run("successful stdout", func(t *testing.T) {
		stdout := stdoutMock{}

		target := "G1 X1 Y3"
		expect := "G1 X-2.000 Y3\n"

		cmd := NewXYCommand()
		cmd.Flags().Lookup("ratio").Value.Set("1")
		cmd.Flags().Lookup("ratio").Changed = true
		cmd.SetIn(strings.NewReader(target))
		cmd.SetOut(&stdout)

		err := cmd.RunE(cmd, []string{})
		if err != nil {
			t.Errorf("expect error nil, got error %v: ", err)
			return
		}

		if stdout.content != expect {
			t.Errorf("expect '%s', got '%s': ", expect, stdout.content)
			return
		}
	})

	t.Run("successful savefile", func(t *testing.T) {
		target := "G1 X1 Y3"
		testFilePath := "../../../data/testsave.gcode"

		cmd := NewXYCommand()
		cmd.Flags().Lookup("ratio").Value.Set("1")
		cmd.Flags().Lookup("ratio").Changed = true
		cmd.Flags().Lookup("output").Value.Set(testFilePath)
		cmd.Flags().Lookup("output").Changed = true
		cmd.SetIn(strings.NewReader(target))

		err := cmd.RunE(cmd, []string{})
		if err != nil {
			t.Errorf("expect error nil, got error %v: ", err)
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

	t.Run("missing savefile", func(t *testing.T) {
		target := "G1 X1 Y3"
		testFilePath := "../../../data/bad folder/testsave.gcode"

		cmd := NewXYCommand()
		cmd.Flags().Lookup("ratio").Value.Set("1")
		cmd.Flags().Lookup("ratio").Changed = true
		cmd.Flags().Lookup("output").Value.Set(testFilePath)
		cmd.Flags().Lookup("output").Changed = true
		cmd.SetIn(strings.NewReader(target))

		err := cmd.RunE(cmd, []string{})
		if err == nil {
			t.Errorf("expect an error, got error nil")
			return
		}
	})

	t.Run("missing flag ratio", func(t *testing.T) {
		target := "G1 X1 Y3"

		cmd := NewXYCommand()
		cmd.SetIn(strings.NewReader(target))

		err := cmd.RunE(cmd, []string{})
		if err == nil {
			t.Errorf("expect an error, got error nil")
			return
		}
	})

	t.Run("missing input", func(t *testing.T) {
		target := "G1 X1 Y3"

		cmd := NewXYCommand()
		cmd.SetIn(strings.NewReader(target))
		cmd.Flags().Lookup("ratio").Value.Set("1")
		cmd.Flags().Lookup("ratio").Changed = true

		err := cmd.RunE(cmd, []string{"badfile"})
		if err == nil {
			t.Errorf("expect an error, got error nil")
			return
		}
	})

	t.Run("missing gcodefile reader", func(t *testing.T) {
		target := "#$*X3"

		cmd := NewXYCommand()
		cmd.SetIn(strings.NewReader(target))
		cmd.Flags().Lookup("ratio").Value.Set("1")
		cmd.Flags().Lookup("ratio").Changed = true

		err := cmd.RunE(cmd, []string{})
		if err == nil {
			t.Errorf("expect an error, got error nil")
			return
		}
	})

	t.Run("missing gcodefile skewer", func(t *testing.T) {
		target := "G1 X3"

		cmd := NewXYCommand()
		cmd.SetIn(strings.NewReader(target))
		cmd.Flags().Lookup("ratio").Value.Set("1")
		cmd.Flags().Lookup("ratio").Changed = true

		err := cmd.RunE(cmd, []string{})
		if err == nil {
			t.Errorf("expect an error, got error nil")
			return
		}
	})
}

func TestGetInput(t *testing.T) {
	testFilePath := "../../../data/test.gcode"
	t.Run("file found", func(t *testing.T) {
		input := strings.NewReader("some content")
		_, err := getInput(input, []string{testFilePath})
		if err != nil {
			t.Errorf("expect not error, got error: %v", err)
		}
	})

	t.Run("file not found", func(t *testing.T) {
		input := strings.NewReader("some content")
		_, err := getInput(input, []string{"../tmp/bad file"})
		if err == nil {
			t.Errorf("expect an error, got error nil")
		}
	})

	t.Run("wihtout file", func(t *testing.T) {
		input := strings.NewReader("some content")
		stream, err := getInput(input, []string{})
		if err != nil {
			t.Errorf("expect not error, got error: %v", err)
		}

		content, err := io.ReadAll(stream)
		if err != nil {
			t.Errorf("expect not error, got error: %v", err)
		}

		if string(content) != "some content" {
			t.Errorf("expect 'some content', got '%s'", string(content))
		}
	})
}
