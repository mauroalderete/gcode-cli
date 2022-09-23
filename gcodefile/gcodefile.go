package gcodefile

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mauroalderete/gcode-core/block/gcodeblock"
)

type GcodeFile struct {
	source io.Reader
	gcodes []*gcodeblock.GcodeBlock
}

func (gf *GcodeFile) Gcodes() []*gcodeblock.GcodeBlock {
	return gf.gcodes
}

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

func (gf *GcodeFile) Update() (io.Reader, error) {
	return nil, nil
}

func (gf *GcodeFile) SaveFile(path string) error {
	return nil
}

func NewFromFile(path string) (*GcodeFile, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed open file in '%s': %v", path, err)
	}

	return NewFromReader(file)
}

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
