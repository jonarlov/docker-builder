package lib

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type consumer func(Config)

// ForEach iterates a list of Config structs, executing the function sent as argument on each struct
func ForEach(l *list.List, fn consumer) {

	for e := l.Front(); e != nil; e = e.Next() {

		fn(e.Value.(Config))
	}
}

// BuildImage builds the image described by the Config struct
func BuildImage(c Config) {

	fmt.Printf("%s in %s\n", c.Dockertag, c.Path)

	println("docker build -t " + c.Dockertag + " " + c.Path)

	cmd := exec.Command("docker", "build", "-t", c.Dockertag, c.Path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// PrintImageList prints the content of the Config struct given as argument
func PrintImageList(c Config) {

	fmt.Printf("Image: %s @ %s\n", c.Dockertag, c.Path)
}
