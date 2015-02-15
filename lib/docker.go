package lib

import (
	"container/list"
	"fmt"
	"strings"
)

// DockerList will list all images in the build chain
func DockerList(cmd Cmd) (out string) {

	// no way to set default values for struct fields, so here we go... This is mainly set to satisfy the tests where we have to call the methods in dobu directly and therefor not using kingpin
	if cmd.Command == "" {
		cmd.Command = "list"
	}

	if cmd.Filename == "" {
		cmd.Filename = "dobu.yml"
	}

	list, out := ReadDobuYamlFiles(cmd)

	fmt.Errorf("%s", list)


	if list != nil && list.Len() > 0 {
		out = "These images will be built:\n"
		fmt.Print(out)
		out += forEach(list, printImageList)
	}

	return
}

// DockerBuild will build all images in the build chain recursively
func DockerBuild(cmd Cmd) (out string) {

	list, out := ReadDobuYamlFiles(cmd)

	if list != nil && list.Len() > 0 {

		out += forEach(list, buildImage)

		out += "Done building you Docker images\n"
		fmt.Print(out)
	}

	return
}

// DockerStop stops all running Docker containers
func DockerStop(time string) (out string) {

	// get all docker container hashes
	hash := dockerPs()

	if len(hash) > 0 {

		// add hashes to docker arguments
		args := append([]string{"stop", "-t", time}, hash...)

		// use ExecOuput to silence docker
		result := ExecOutput("docker", args)
		out = result.Out()
	} else {
		out = "No Docker containers to stop\n"
		fmt.Print(out)
	}

	return
}

//DockerDeleteContainers deletes all Docker containers
func DockerDeleteContainers() (out string) {

	// get all docker container hashes
	hash := dockerPs()

	if len(hash) > 0 {

		// add hashes to docker arguments
		args := append([]string{"rm", "-f"}, hash...)

		result := ExecCommand("docker", args)
		out = result.Out()
	} else {
		out = "No Docker containers to delete\n"
		fmt.Print(out)
	}

	return
}

//DockerDeleteImages deletes all Docker images
func DockerDeleteImages() (out string) {

	hash := dockerImages()

	if len(hash) > 0 {
		// append hashes to docker arguments
		args := append([]string{"rmi", "-f"}, hash...)

		result := ExecCommand("docker", args)
		out = result.Out()
	} else {
		out = "No Docker images to delete\n"
		fmt.Print(out)
	}

	return
}

// dockerPs returns an array of hashes for all containers
func dockerPs() (hash []string) {

	// get hash of all running containers
	result := ExecOutput("docker", []string{"ps", "-a", "-q"})
	result.ErrorHandling()

	if len(result.stdout) > 0 {

		// split docker hash output on newline
		hash = strings.Split(string(result.stdout), "\n")
	}

	return hash
}

// dockerImages returns an array of hashes for all images
func dockerImages() (hash []string) {

	// get hash of all running containers
	result := ExecOutput("docker", []string{"images", "-q"})
	result.ErrorHandling()

	if len(result.stdout) > 0 {

		// split docker hash output on newline
		hash = strings.Split(string(result.stdout), "\n")
	}

	return hash
}

type consumer func(Config) (out string)

// forEach iterates a list of Config structs, executing the function sent as argument on each struct
func forEach(l *list.List, consumer consumer) (out string) {

	for e := l.Front(); e != nil; e = e.Next() {

		out += consumer(e.Value.(Config))
	}

	return
}

// buildImage builds the image described by the Config struct
func buildImage(c Config) (out string) {

	out = fmt.Sprintf("Building %s in %s\n", c.Dockertag, c.Path)
	fmt.Print(out)

	result := ExecCommand("docker", []string{"build", "-t", c.Dockertag, c.Path})

	out += result.Out()

	return
}

// printImageList prints the content of the Config struct given as argument
func printImageList(c Config) (out string) {

	out = fmt.Sprintf("Image: %s @ %s\n", c.Dockertag, c.Path)
	fmt.Print(out)

	return
}
