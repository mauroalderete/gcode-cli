package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	rootCmd := NewRootCommand("v1.0.0")
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("want not error, got %v", err)
		return
	}

	out, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Errorf("want readAll buffer, got %v", err)
		return
	}

	if len(out) == 0 {
		t.Error("want out not empty, got out empty")
	}
}

func Test_Execute(t *testing.T) {
	rootCmd := NewRootCommand("v1.0.0")
	bufferStdout := new(bytes.Buffer)
	bufferStderr := new(bytes.Buffer)
	rootCmd.SetOut(bufferStdout)
	rootCmd.SetErr(bufferStderr)

	pathedExit := struct {
		Called     bool
		CalledWith int
	}{
		Called:     false,
		CalledWith: 0,
	}

	exit := func(code int) {
		pathedExit.Called = true
		pathedExit.CalledWith = code
	}

	Execute(rootCmd, exit)

	if pathedExit.Called {
		t.Error("expected terminator to be called, but not happen")
		return
	}
}

func Test_ExecuteWithTerminator(t *testing.T) {
	rootCmd := NewRootCommand("v1.0.0")
	bufferStdout := new(bytes.Buffer)
	bufferStderr := new(bytes.Buffer)
	rootCmd.SetOut(bufferStdout)
	rootCmd.SetErr(bufferStderr)
	rootCmd.SetArgs([]string{"some", "bad", "args"})

	pathedExit := struct {
		Called     bool
		CalledWith int
	}{
		Called:     false,
		CalledWith: 0,
	}

	exit := func(code int) {
		pathedExit.Called = true
		pathedExit.CalledWith = code
	}

	Execute(rootCmd, exit)

	if !pathedExit.Called {
		t.Error("expected terminator to be called, but not happen")
		return
	}

	if pathedExit.CalledWith != 1 {
		t.Errorf("want terminator to be called with code 1, got %v", pathedExit.CalledWith)
		return
	}
}
