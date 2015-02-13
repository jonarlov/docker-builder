package lib

import (
	"container/list"
	"fmt"
)

type consumer func(Config) (err error)

// ForEach iterates a list of Config structs, executing the function sent as argument on each struct
func ForEach(l *list.List, consumer consumer) (err error) {

	for e := l.Front(); e != nil; e = e.Next() {

		err = consumer(e.Value.(Config))

		if err != nil {
			return err
		}
	}

	return
}

// BuildImage builds the image described by the Config struct
func BuildImage(c Config) (err error) {

	fmt.Printf("%s in %s\n", c.Dockertag, c.Path)

	println("docker build -t " + c.Dockertag + " " + c.Path)

	return ExecCommand("docker", "build", "-t", c.Dockertag, c.Path)
}

// PrintImageList prints the content of the Config struct given as argument
func PrintImageList(c Config) (err error) {

	_, err = fmt.Printf("Image: %s @ %s\n", c.Dockertag, c.Path)
	return
}
