package lib

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// ReadDobuYamlFiles reads all yaml files recursively and parse them into our Config struct
func ReadDobuYamlFiles(cmd Cmd) (*list.List, error) {

	list := list.New()

	path, _ := filepath.Abs(cmd.Path)
	err := recursiveReadFiles(path, cmd, list)

	return list, err
}

func recursiveReadFiles(path string, cmd Cmd, l *list.List) error {

	dobuContent, err := readFile(path + "/" + cmd.Filename)
	if err != nil {
		return fmt.Errorf("No %s in %s...", cmd.Filename, path)
	}

	if _, err := os.Stat(path + "/Dockerfile"); err != nil {
		return fmt.Errorf("No Dockerfile in %s...", path)
	}

	var conf Config
	unmarshal(dobuContent, &conf)

	abs, _ := filepath.Abs(path)
	conf.Path = abs

	if conf.Dockertag == "" {
		return fmt.Errorf("dockertag is required in %s/%s", conf.Path, cmd.Filename)
	}

	l.PushFront(conf)

	if conf.Parent != "" {
		return recursiveReadFiles(path+"/"+conf.Parent, cmd, l)
	}

	return nil
}

func readFile(s string) ([]byte, error) {

	result, err := ioutil.ReadFile(s)
	return result, err
}

func unmarshal(in []byte, conf *Config) {
	yaml.Unmarshal(in, conf)
}
