package main

import (
	"fmt"
	"os"

	"github.com/orby/docker-builder/lib"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	// dobu
	app = kingpin.New("dobu", "Dobu is a recursive Docker image builder.").Version("0.2.1")

	// dobu list
	listCommand      = kingpin.Command("list", "List images in the build chain")
	listFilenameFlag = listCommand.Flag("file", "Alternate dobu.yml filename").Default("dobu.yml").Short('f').String()
	listWdFlag       = listCommand.Flag("working-directory", "Change working directory").Default(".").Short('w').String()

	// dobu build
	buildCommand      = kingpin.Command("build", "Build images in the build chain recursivly")
	buildFilenameFlag = buildCommand.Flag("file", "Alternate dobu.yml filename").Default("dobu.yml").Short('f').String()
	buildWdFlag       = buildCommand.Flag("working-directory", "Change working directory").Default(".").Short('w').String()

	// dobu stop
	stopCommand  = kingpin.Command("stop", "Stop all running Docker containers by sending SIGTERM and then SIGKILL after a grace period")
	stopTimeFlag = stopCommand.Flag("time", "Number of seconds to wait for the container to stop before killing it. Default is 10 seconds").Default("10").Short('t').String()

	// dobu rm
	deleteCommand   = kingpin.Command("delete", "Delete all containers, or all images or both containers and images")
	containerDelete = deleteCommand.Command("containers", "Delete all Docker containers")
	imageDelete     = deleteCommand.Command("images", "Delete all Docker images")
	allDelete       = deleteCommand.Command("all", "Delete all Docker containers and Docker images")
)

func main() {

	switch kingpin.Parse() {

	case "list":
		doList(lib.Cmd{Path: *listWdFlag, Filename: *listFilenameFlag, Command: "list"})

	case "build":
		doBuild(lib.Cmd{Path: *buildWdFlag, Filename: *buildFilenameFlag, Command: "build"})

	case "stop":
		doStop(*stopTimeFlag)

	case "delete containers":
		doDelete("containers")

	case "delete images":
		doDelete("images")

	case "delete all":
		doDelete("all")

	default:
		usage()
	}
}

func doList(cmd lib.Cmd) (out string) {

	return lib.DockerList(cmd)
}

func doBuild(cmd lib.Cmd) (out string) {

	return lib.DockerBuild(cmd)
}

func doStop(time string) (out string) {

	return lib.DockerStop(time)
}

func doDelete(arg string) (out string) {

	switch arg {

	case "containers":
		out = lib.DockerDeleteContainers()
	case "images":
		out = lib.DockerDeleteImages()
	case "all":
		out1 := lib.DockerDeleteContainers()
		out2 := lib.DockerDeleteImages()
		out = out1 + "\n" + out2
	case "":
		out = "You need to specify either \"delete containers\", \"delete images\" or \"delete all\" when calling delete. See \"dobu help delete\" for more information"
		fmt.Println(out)
	default:
		out := "Unknown delete directive: " + arg
		fmt.Println(out)
	}

	return
}

func usage() {
	app.Usage(os.Stdout)
}
