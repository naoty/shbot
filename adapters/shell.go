package adapters

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Shell is an adapter which read messages from shell and write messages to shell.
type Shell struct {
	input       io.Reader
	output      io.Writer
	errorOutput io.Writer
}

// NewShell returns a new Shell.
func NewShell() *Shell {
	return &Shell{
		input:       os.Stdin,
		output:      os.Stdout,
		errorOutput: os.Stdout,
	}
}

// Prepare runs preparation before accepting messages.
func (adapter *Shell) Prepare() {
	fmt.Fprint(adapter.output, "shbot> ")
}

// ReadMessage returns messages from input.
func (adapter *Shell) ReadMessage() string {
	reader := bufio.NewReader(adapter.input)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(adapter.errorOutput)
	}

	message := strings.TrimRight(line, "\n")
	return message
}

// WriteMessage send a given message to output.
func (adapter *Shell) WriteMessage(message string) {
	fmt.Fprintln(adapter.output, message)
}

// WriteError send an error message to output.
func (adapter *Shell) WriteError(err error) {
	fmt.Fprintln(adapter.errorOutput, err)
}
