package main

import (
	"fmt"
	"github.com/orby/dobu/lib"
	"gopkg.in/alecthomas/kingpin.v1"
	"log"
	"os"
	//"strings"
)

var (
	app          = kingpin.New("dobu", "A Docker image builder.")
	listCommand  = app.Command("list", "List docker images that would be build")
	buildCommand = app.Command("build", "Build docker images recursivly")
	wdFlag       = app.Flag("working-directory", "If you want to change working directory").Default(".").Short('w').String()
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "list":
		list(*wdFlag)
	case "build":
		build(*wdFlag)
	}
}

func list(path string) {
	fmt.Println("These images will be built:")

	list, err := lib.ReadDobuYamlFiles(path)

	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
		return
	}

	lib.ForEach(list, lib.PrintImageList)
}

func build(path string) {
	fmt.Println("Inside build")

	list, err := lib.ReadDobuYamlFiles(path)

	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
		return
	}

	lib.ForEach(list, lib.BuildImage)
}
