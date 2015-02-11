package lib

import (
	"container/list"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Parent    string
	Dockertag string
	Path      string
}

func ReadDobuYamlFiles() (*list.List, error) {

	list := list.New()

	path, _ := filepath.Abs(".")
	err := recursiveReadFiles(path, list)

	return list, err
}

func recursiveReadFiles(path string, l *list.List) error {

	//return errors.New("PIO")
	dobuContent, err := readFile(path + "/dobu.yml")
	if err != nil {
		return fmt.Errorf("No dobu.yml in %s...", path)
	}

	// //TODO: Check for Dockerfile
	if _, err := os.Stat(path + "/Dockerfile"); err != nil {
		return fmt.Errorf("No Dockerfile in %s...", path)
	}

	var conf Config
	unmarshal(dobuContent, &conf)

	abs, err := filepath.Abs(path)
	conf.Path = abs

	if conf.Dockertag == "" {
		return fmt.Errorf("dockertag is required in %s/dobu.yml", conf.Path)
	}

	l.PushFront(conf)

	if conf.Parent != "" {
		return recursiveReadFiles(path+"/"+conf.Parent, l)
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
