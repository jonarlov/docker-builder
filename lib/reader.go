package lib

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config holding values from the yaml files
type Config struct {
	Parent    string
	Dockertag string
	Path      string
}

// ReadDobuYamlFiles reads all yaml files recursively and parse them into our Config struct
func ReadDobuYamlFiles(path string, filename string) (*list.List, error) {

	list := list.New()

	path, _ = filepath.Abs(path)
	err := recursiveReadFiles(path, filename, list)

	return list, err
}

func recursiveReadFiles(path string, filename string, l *list.List) error {

	dobuContent, err := readFile(path + "/" + filename)
	if err != nil {
		return fmt.Errorf("No %s in %s...", filename, path)
	}

	if _, err := os.Stat(path + "/Dockerfile"); err != nil {
		return fmt.Errorf("No Dockerfile in %s...", path)
	}

	var conf Config
	unmarshal(dobuContent, &conf)

	abs, _ := filepath.Abs(path)
	conf.Path = abs

	if conf.Dockertag == "" {
		return fmt.Errorf("dockertag is required in %s/%s", conf.Path, filename)
	}

	l.PushFront(conf)

	if conf.Parent != "" {
		return recursiveReadFiles(path+"/"+conf.Parent, filename, l)
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
