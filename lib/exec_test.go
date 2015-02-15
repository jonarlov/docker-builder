package lib

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// if any setup
//func TestMain(m *testing.M) { os.Exit(m.Run()) }

//can't get ExecCommand to return a copy of stdout, yet
func TestExecCommand(t *testing.T) {

	echo := "Testing output"

	result := ExecCommand("echo", strings.Fields(echo))

	// should output echo + newline
	assert.Equal(t, result.Out(), echo+"\n")
}

func TestExecOutput(t *testing.T) {

	echo := "Testing output"

	result := ExecOutput("echo", strings.Fields(echo))

	// should output echo + newline
	assert.Equal(t, result.Out(), echo+"\n")
}

func TestExecOutputErrorHandling(t *testing.T) {

	// testing error handling when calling a non existing application
	result := ExecOutput("golangcommandnotavailable", strings.Fields("-test"))

	assert.Contains(t, result.err.Error(), "golangcommandnotavailable")
}
