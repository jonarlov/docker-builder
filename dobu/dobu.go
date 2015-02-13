package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/orby/docker-builder/lib"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	app          = kingpin.New("dobu", "Dobu is a recursive Docker image builder.").Version("0.2.1")
	listCommand  = app.Command("list", "List images in the build chain")
	buildCommand = app.Command("build", "Build images in the build chain recursivly")
	stopCommand  = app.Command("stop", "Stop all running containers by sending SIGTERM and then SIGKILL after a grace period")
	stopTimeFlag = stopCommand.Flag("time", "Number of seconds to wait for the container to stop before killing it. Default is 10 seconds").Default("10").Short('t').String()
	wdFlag       = app.Flag("working-directory", "Change working directory").Default(".").Short('w').String()
	filenameFlag = app.Flag("file", "Alternate dobu.yml filename").Default("dobu.yml").Short('f').String()
)

func main() {

	p := kingpin.MustParse(app.Parse(os.Args[1:]))

	cmd := lib.Cmd{Path: *wdFlag, Filename: *filenameFlag}

	switch p {
	case "list":
		list(cmd)
	case "build":
		build(cmd)
	case "stop":
		stop(*stopTimeFlag)
	default:
		usage()
	}
}

func list(cmd lib.Cmd) {
	fmt.Println("These images will be built:")

	list, err := lib.ReadDobuYamlFiles(cmd)
	exitIfError(100, err, "")

	err = lib.ForEach(list, lib.PrintImageList)
	exitIfError(101, err, "")
}

func build(cmd lib.Cmd) {
	fmt.Println("Building:")

	list, err := lib.ReadDobuYamlFiles(cmd)
	exitIfError(200, err, "")

	err = lib.ForEach(list, lib.BuildImage)
	exitIfError(201, err, "")

	fmt.Println("Done building you Docker images")
}

func stop(time string) {

	fmt.Printf("docker stop -t %s $(docker ps -a -q)\n", time)

	// get hash of all running containers
	out, err := lib.ExecOutput("docker", "ps", "-a", "-q")
	exitIfError(301, err, string(out))

	// split docker hash output on newline
	hash := strings.Split(string(out), "\n")

	// docker arguments
	args := []string{"stop", "-t", time}

	// append hash to docker arguments
	cmd := append(args, hash...)

	err = lib.ExecCommand("docker", cmd...)
	exitIfError(300, err, "")
}

func usage() {

	app.Usage(os.Stdout)
}

func exitIfError(code int, err error, message string) {

	if err != nil {
		fmt.Printf("ERROR: %s\n", message) // TODO: error code
		os.Exit(code)
	}
}
