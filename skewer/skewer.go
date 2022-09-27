package skewer

import (
	"fmt"

	"github.com/mauroalderete/gcode-core/block/gcodeblock"
	"github.com/mauroalderete/gcode-core/gcode"
	"github.com/mauroalderete/gcode-core/gcode/addressablegcode"
)

func SkewXY(ratio float32, blocks []gcodeblock.GcodeBlock) ([]gcodeblock.GcodeBlock, error) {

	// List of all blocks include them that must be transformed
	transformed := make([]gcodeblock.GcodeBlock, 0)

	for i, block := range blocks {

		// Only with G command
		if block.Command().String() != "G1" && block.Command().String() != "G0" {
			transformed = append(transformed, block)
			continue
		}

		//Get parameters to gcode X and Y as addressable float32
		parameters := block.Parameters()

		px, err := getParameterFloat32('X', parameters)
		if err != nil {
			transformed = append(transformed, block)
			continue
		}

		py, err := getParameterFloat32('Y', parameters)
		if err != nil {
			return nil, fmt.Errorf("skewXY expect parameter Y but failed to recover from [%d][%s]: %v", i, block.ToLine("%l %c %p %k %m"), err)
		}

		// Applies transform
		for i := 0; i <= len(parameters); i++ {

			if parameters[i].Word() == 'X' {

				px.SetAddress(px.Address() - py.Address()*ratio)
				parameters[i] = px

				break
			}
		}

		// Generates the new block with parameters transformed
		blockToParse := block.ToLine("%c")
		for _, p := range parameters {
			blockToParse += fmt.Sprintf(" %s", p.String())
		}
		blockToParse += block.ToLine("%k %m")

		blockTransformed, err := gcodeblock.Parse(blockToParse)
		if err != nil {
			return nil, fmt.Errorf("failed to generate a new block transformed from block [%d][%s]: %v", i, block.ToLine("%l %c %p %k %m"), err)
		}

		// Add block transformed in the list
		transformed = append(transformed, *blockTransformed)
	}

	return transformed, nil
}

func getParameterFloat32(parameterWord byte, parameters []gcode.Gcoder) (*addressablegcode.Gcode[float32], error) {
	found := false

	parameterFound, err := addressablegcode.New[float32](parameterWord, 0)
	if err != nil {
		return nil, fmt.Errorf("failed get parameter '%c': %v", parameterWord, err)
	}

	for _, parameter := range parameters {
		if parameter.Word() == parameterWord {
			found = true

			if !parameter.HasAddress() {
				return nil, fmt.Errorf("block bad formed, expected an address available to '%c' parameter", parameterWord)
			}

			switch value := parameter.(type) {
			case gcode.AddressableGcoder[float32]:
				parameterFound.SetAddress(value.Address())
			case gcode.AddressableGcoder[int32]:
				parameterFound.SetAddress(float32(value.Address()))
			case gcode.AddressableGcoder[uint32]:
				parameterFound.SetAddress(float32(value.Address()))
			default:
				return nil, fmt.Errorf("failed get addres from '%c' parameter by assertion", parameterWord)
			}

			break
		}
	}

	if !found {
		return nil, fmt.Errorf("parameter '%c' not found", parameterWord)
	}

	return parameterFound, nil
}

func SkewXZ(ratio float32, blocks []gcodeblock.GcodeBlock) ([]gcodeblock.GcodeBlock, error) {
	return nil, nil
}

func SkewYZ(ratio float32, blocks []gcodeblock.GcodeBlock) ([]gcodeblock.GcodeBlock, error) {
	return nil, nil
}
