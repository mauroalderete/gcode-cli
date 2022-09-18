package description

import (
	"bufio"
	"fmt"

	"github.com/mauroalderete/gcode-core/block/gcodeblock"
)

// Parse implements [Descriptionable.Parse]
func (d *Description) Parse() error {

	blocksCount := 0
	linesCount := 0

	scanner := bufio.NewScanner(d.source)
	for scanner.Scan() {
		linesCount++

		line := scanner.Text()

		_, err := gcodeblock.Parse(line)
		if err == nil {
			blocksCount++
		}
	}

	if linesCount == 0 {
		if d.filename != "" {
			return fmt.Errorf("file %s is empty", d.filename)
		} else {
			return fmt.Errorf("no entry provided on stdin")
		}
	}

	d.linesCount = linesCount
	d.blocksCount = blocksCount

	return nil
}
