package lib

import (
	"container/list"
	"fmt"
)

type consumer func(Config)

// ForEach iterates a list of Config structs, executing the function sent as argument on each struct
func ForEach(l *list.List, consumer consumer) {

	for e := l.Front(); e != nil; e = e.Next() {

		consumer(e.Value.(Config))
	}
}

// BuildImage builds the image described by the Config struct
func BuildImage(c Config) {

	fmt.Printf("%s in %s\n", c.Dockertag, c.Path)

	ExecCommand("docker", "build", "-t", c.Dockertag, c.Path)
}

// PrintImageList prints the content of the Config struct given as argument
func PrintImageList(c Config) {

	fmt.Printf("Image: %s @ %s\n", c.Dockertag, c.Path)
}
