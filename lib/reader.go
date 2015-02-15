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
func ReadDobuYamlFiles(cmd Cmd) (l *list.List, out string) {

	list := list.New()

	path, _ := filepath.Abs(cmd.Path)
	out = recursiveReadFiles(path, cmd, list)

	return list, out
}

func recursiveReadFiles(path string, cmd Cmd, l *list.List) (out string) {

	dobuContent, err := readFile(path + "/" + cmd.Filename)
	if err != nil {
		out = fmt.Sprintf("No %s in %s. See \"dobu help %s\" for using -f to specify another filename...\n", cmd.Filename, path, cmd.Command)
		fmt.Print(out)
		return
	}

	if _, err := os.Stat(path + "/Dockerfile"); err != nil {
		out = fmt.Sprintf("No Dockerfile in %s...\n", path)
		fmt.Print(out)
		return
	}

	var conf Config
	unmarshal(dobuContent, &conf)

	abs, _ := filepath.Abs(path)
	conf.Path = abs

	if conf.Dockertag == "" {
		out = fmt.Sprintf("dockertag is required in %s/%s\n", conf.Path, cmd.Filename)
		fmt.Print(out)
		return
	}

	l.PushFront(conf)

	if conf.Parent != "" {
		recursiveReadFiles(path+"/"+conf.Parent, cmd, l)
	}

	return
}

// read and return file content
func readFile(s string) ([]byte, error) {

	result, err := ioutil.ReadFile(s)
	return result, err
}

// map content from file to our Cmd struct
func unmarshal(in []byte, conf *Config) {
	yaml.Unmarshal(in, conf)
}
