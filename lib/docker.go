package lib

import (
	"container/list"
	"fmt"
	"strings"
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

	ExecCommand("docker", []string{"build", "-t", c.Dockertag, c.Path})
}

// PrintImageList prints the content of the Config struct given as argument
func PrintImageList(c Config) {

	fmt.Printf("Image: %s @ %s\n", c.Dockertag, c.Path)
}

// DockerList will list all images in the build chain
func DockerList(cmd Cmd) {

	list := ReadDobuYamlFiles(cmd)

	if list.Len() > 0 {
		fmt.Println("These images will be built:")
		ForEach(list, PrintImageList)
	}
}

// DockerBuild will build all images in the build chain recursively
func DockerBuild(cmd Cmd) {

	list := ReadDobuYamlFiles(cmd)

	if list.Len() > 0 {
		fmt.Println("Building:")

		ForEach(list, BuildImage)

		fmt.Println("Done building you Docker images")
	}
}

// DockerStop stops all running Docker containers
func DockerStop(time string) {

		// get all docker container hashes
		hash := dockerPs()

		if len(hash) > 0 {

			// add hashes to docker arguments
			args := append([]string{"stop", "-t", time}, hash...)

			// use ExecOuput to silence docker
			ExecOutput("docker", args)
		} else {
			fmt.Println("No Docker containers to stop")
		}
}

//DockerDeleteContainers deletes all Docker containers
func DockerDeleteContainers() {

	// get all docker container hashes
	hash := dockerPs()

	if len(hash) > 0 {

			// add hashes to docker arguments
			args := append([]string{"rm", "-f"}, hash...)

			ExecCommand("docker", args)
	} else {
		fmt.Println("No Docker containers to delete")
	}

}

//DockerDeleteImages deletes all Docker images
func DockerDeleteImages() {

	hash := dockerImages()

	if len(hash) > 0 {
		// append hashes to docker arguments
		args := append([]string{"rmi", "-f"}, hash...)

		ExecCommand("docker", args)
	} else {
		fmt.Println("No Docker images to delete")
	}


}

// dockerPs returns an array of hashes for all containers
func dockerPs() (hash []string) {

	// get hash of all running containers
	out := ExecOutput("docker", []string{"ps", "-a", "-q"})

	if len(out) > 0 {

		// split docker hash output on newline
		hash = strings.Split(string(out), "\n")
	}

	return hash
}

// dockerImages returns an array of hashes for all images
func dockerImages() (hash []string) {

	// get hash of all running containers
	out := ExecOutput("docker", []string{"images", "-q"})

	if len(out) > 0 {

		// split docker hash output on newline
		hash = strings.Split(string(out), "\n")
	}

	return hash
}
