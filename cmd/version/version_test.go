package version

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
)

func Test_ExecuteCommand(t *testing.T) {
	rootCmdMock := &cobra.Command{
		Use:     "gcode-cli",
		Version: "1.0.0",
	}
	expected := fmt.Sprintf("%s version %s\n", rootCmdMock.Name(), rootCmdMock.Version)

	rootCmdMock.SetArgs([]string{"version"})
	buffer := new(bytes.Buffer)
	rootCmdMock.SetOut(buffer)

	versionCmd := NewVersionCommand()
	rootCmdMock.AddCommand(versionCmd)

	versionCmd.Execute()
	out, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != expected {
		t.Errorf("want '%s', got '%s", expected, string(out))
	}
}
