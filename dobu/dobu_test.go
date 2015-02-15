package main

import (
  "testing"
  "fmt"
  "path/filepath"
  "github.com/orby/docker-builder/lib"

  "github.com/stretchr/testify/assert"
)

// if any setup
//func TestMain(m *testing.M) { os.Exit(m.Run()) }

func TestSuccessfulListCommand(t *testing.T) {

  output := doList(lib.Cmd{Path: "../test/base-images/debian-wheezy", Filename: "dobu.yml"})
  assert.Contains(t, output, "These images will be built:\nImage: company/debian:7 @ /") // the end / will probably fail on windows...
}

func TestListOutputWithNoCommandAndNoFilename(t *testing.T) {

  output := doList(lib.Cmd{Path: "../test/base-images/debian-wheezy"})
  assert.Contains(t, output, "These images will be built:\nImage: company/debian:7 @ /") // the end / will probably fail on windows...
}

func TestListOutputWithNoDockerfileInDirectory(t *testing.T) {

  output := doList(lib.Cmd{Path: "../test", Filename: "test.yml"})

  // find path that will be used in output, can't really hardcode path from my computer here
  path, _ := filepath.Abs("../test")

  assert.Equal(t, output, fmt.Sprintf("No Dockerfile in %s...\n", path))
}

func TestListOutputWithUnknownFilename(t *testing.T) {

  output := doList(lib.Cmd{Filename: "vif.yml", Command: "list"})
  assert.Equal(t, output, "No vif.yml in /Users/orby/src/go/src/github.com/orby/docker-builder/dobu. See \"dobu help list\" for using -f to specify another filename...\n")
}
