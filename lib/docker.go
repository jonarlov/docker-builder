package lib

import (
	"container/list"
	"fmt"
	//"os/exec"
)

type consumer func(Config)

func ForEach(l *list.List, fn consumer) {

	for e := l.Front(); e != nil; e = e.Next() {

		fn(e.Value.(Config))
	}
}

func BuildImage(c Config) {

	fmt.Printf("Building Docker image located in %s\n", c.Path)

	println("docker build -t " + c.Dockertag + " " + c.Path)

}

func PrintImageList(c Config) {

	fmt.Printf("Image: %s @ %s\n", c.Dockertag, c.Path)
}
