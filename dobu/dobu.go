package main

import (
	"fmt"
	"log"
	"os"

	"github.com/orby/docker-builder/lib"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	app          = kingpin.New("dobu", "A Docker image builder.").Version("0.1.1")
	listCommand  = app.Command("list", "List docker images that would be build")
	buildCommand = app.Command("build", "Build docker images recursivly")
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
	default:
		usage()
	}
}

func list(cmd lib.Cmd) {
	fmt.Println("These images will be built:")

	list, err := lib.ReadDobuYamlFiles(cmd)

	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
		return
	}

	lib.ForEach(list, lib.PrintImageList)
}

func build(cmd lib.Cmd) {
	fmt.Println("Inside build")

	list, err := lib.ReadDobuYamlFiles(cmd)

	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
		return
	}

	lib.ForEach(list, lib.BuildImage)
}

func usage() {

	app.Usage(os.Stdout)
}
