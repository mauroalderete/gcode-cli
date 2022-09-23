package gcodefile

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mauroalderete/gcode-core/block/gcodeblock"
)

// GcodeFile model a file or another type source that contain a 3D object written in gcode format.
//
// GcodeFile implements basic methods to handle the file, open, read it and save.
// Exposes the list of Gcodes that define the 3d object to process the model or apply changes.
// Series of methods allows Update the source of model and save it as a file.
type GcodeFile struct {
	source io.Reader
	gcodes []*gcodeblock.GcodeBlock
}

// Gcodes returns the lis of blocks found in the source, after parse it, as a list of [gcodeblock.GcodeBlock]
func (gf *GcodeFile) Gcodes() []*gcodeblock.GcodeBlock {
	return gf.gcodes
}

// Source returns the source as [io.Reader] that store the blocks written as plain text line by line.
func (gf *GcodeFile) Source() io.Reader {
	return gf.source
}

// Refresh parse the gcode source loaded to get a list of [gcodeblock.GcodeBlock]
//
// Read the source content line by line to get his [gcodeblock.GcodeBlock] representation.
//
// If the source is empty or only contains comments then Refresh parse a list empty.
//
// If the source is nil then refresh returns an error.
//
// If the source contain an invalid block then Refresh returns an error.
func (gf *GcodeFile) Refresh() error {

	scanner := bufio.NewScanner(gf.source)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		//line is empty
		if len(line) == 0 {
			continue
		}

		//line is a comment
		if strings.HasPrefix(line, ";") {
			continue
		}

		block, err := gcodeblock.Parse(line)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		gf.gcodes = append(gf.gcodes, block)
	}

	return nil
}

// Update takes the last state of the blocks in the list stored and updates the source with the changes made to this moment.
func (gf *GcodeFile) Update() error {
	return nil
}

func (gf *GcodeFile) SaveFile(path string) error {
	return nil
}

// NewFromFile returns a new instance of [GcodeFile] from a file with the gcode blocks contented.
func NewFromFile(path string) (*GcodeFile, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed open file in '%s': %v", path, err)
	}

	return NewFromReader(file)
}

// NewFromReader returns a new instance of [GcodeFile] from an [io.Reader] object with the gcode blocks contented.
func NewFromReader(source io.Reader) (*GcodeFile, error) {

	gf := GcodeFile{
		gcodes: make([]*gcodeblock.GcodeBlock, 0),
		source: source,
	}

	err := gf.Refresh()
	if err != nil {
		return nil, fmt.Errorf("failed parse source: %v", err)
	}

	return &gf, nil
}
