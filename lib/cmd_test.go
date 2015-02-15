package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// if any setup
//func TestMain(m *testing.M) { os.Exit(m.Run()) }

func TestCmdNamedFields(t *testing.T) {

	var cmd Cmd

	assert.Empty(t, cmd.Path)
	assert.Empty(t, cmd.Filename)
	assert.Empty(t, cmd.Command)
}

func TestCmdValues(t *testing.T) {

	cmd := Cmd{Path: "../", Filename: "dobu.yml", Command: "test"}

	assert.Equal(t, cmd.Path, "../")
	assert.Equal(t, cmd.Filename, "dobu.yml")
	assert.Equal(t, cmd.Command, "test")
}
