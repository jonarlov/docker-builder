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
func ReadDobuYamlFiles(cmd Cmd) (*list.List) {

	list := list.New()

	path, _ := filepath.Abs(cmd.Path)
	recursiveReadFiles(path, cmd, list)

	return list
}

func recursiveReadFiles(path string, cmd Cmd, l *list.List) {

	dobuContent, err := readFile(path + "/" + cmd.Filename)
	if err != nil {
		fmt.Printf("No %s in %s. See \"dobu help %s\" -f to use another filename...\n", cmd.Filename, path, cmd.Command)
		os.Exit(1)
	}

	if _, err := os.Stat(path + "/Dockerfile"); err != nil {
		fmt.Printf("No Dockerfile in %s...\n", path)
		os.Exit(1)
	}

	var conf Config
	unmarshal(dobuContent, &conf)

	abs, _ := filepath.Abs(path)
	conf.Path = abs

	if conf.Dockertag == "" {
		fmt.Printf("dockertag is required in %s/%s\n", conf.Path, cmd.Filename)
		os.Exit(1)
	}

	l.PushFront(conf)

	if conf.Parent != "" {
		recursiveReadFiles(path+"/"+conf.Parent, cmd, l)
	}
}

func readFile(s string) ([]byte, error) {

	result, err := ioutil.ReadFile(s)
	return result, err
}

func unmarshal(in []byte, conf *Config) {
	yaml.Unmarshal(in, conf)
}
