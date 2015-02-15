package lib

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

// if any setup
//func TestMain(m *testing.M) { os.Exit(m.Run()) }

func TestConfigNamedFields(t *testing.T) {

  var config Config

  assert.Empty(t, config.Parent)
  assert.Empty(t, config.Dockertag)
  assert.Empty(t, config.Path)
}

func TestConfigValues(t *testing.T) {

  config := Config{Parent: "../", Dockertag: "dobu.yml", Path: "test"}

  assert.Equal(t, config.Parent, "../")
  assert.Equal(t, config.Dockertag, "dobu.yml")
  assert.Equal(t, config.Path, "test")
}
