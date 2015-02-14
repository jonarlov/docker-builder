package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/orby/docker-builder/lib"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	// dobu
	app = kingpin.New("dobu", "Dobu is a recursive Docker image builder.").Version("0.2.1")

	// dobu list
	listCommand      = app.Command("list", "List images in the build chain")
	listFilenameFlag = listCommand.Flag("file", "Alternate dobu.yml filename").Default("dobu.yml").Short('f').String()
	listWdFlag       = listCommand.Flag("working-directory", "Change working directory").Default(".").Short('w').String()

	// dobu build
	buildCommand      = app.Command("build", "Build images in the build chain recursivly")
	buildFilenameFlag = buildCommand.Flag("file", "Alternate dobu.yml filename").Default("dobu.yml").Short('f').String()
	buildWdFlag       = buildCommand.Flag("working-directory", "Change working directory").Default(".").Short('w').String()

	// dobu stop
	stopCommand  = app.Command("stop", "Stop all running Docker containers by sending SIGTERM and then SIGKILL after a grace period")
	stopTimeFlag = stopCommand.Flag("time", "Number of seconds to wait for the container to stop before killing it. Default is 10 seconds").Default("10").Short('t').String()
)

func main() {

	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch command {
	case "list":
		list(lib.Cmd{Path: *listWdFlag, Filename: *listFilenameFlag, Command: command})
	case "build":
		build(lib.Cmd{Path: *buildWdFlag, Filename: *buildFilenameFlag, Command: command})
	case "stop":
		stop(*stopTimeFlag)
	default:
		usage()
	}
}

func list(cmd lib.Cmd) {

	list := lib.ReadDobuYamlFiles(cmd)

	if list.Len() > 0 {
		fmt.Println("These images will be built:")
		lib.ForEach(list, lib.PrintImageList)
	}
}

func build(cmd lib.Cmd) {

	list := lib.ReadDobuYamlFiles(cmd)

	if list.Len() > 0 {
		fmt.Println("Building:")

		lib.ForEach(list, lib.BuildImage)

		fmt.Println("Done building you Docker images")
	}
}

func stop(time string) {

	fmt.Printf("docker stop -t %s $(docker ps -a -q)\n", time)

	// get hash of all running containers
	out := lib.ExecOutput("docker", "ps", "-a", "-q")

	// split docker hash output on newline
	hash := strings.Split(string(out), "\n")

	// docker arguments array
	args := []string{"stop", "-t", time}

	// append hash to docker arguments
	cmd := append(args, hash...)

	lib.ExecCommand("docker", cmd...)
}

func usage() {

	app.Usage(os.Stdout)
}
