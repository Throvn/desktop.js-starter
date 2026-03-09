package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	flag "github.com/spf13/pflag"
)

//go:embed version
var version string

func main() {

	if runtime.GOOS != "darwin" {
		Info("Desktop.js is currently only available for macOS.\nYou are not running on macOS.\nContribute to make it available on more platforms!")
	}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, `
init <location>        | Creates a new project at the specified location
watch <location>       | Starts live reload for the given project
bundle <app_name>      | (MacOS only) Creates an appbundle
version                | Prints the current starter version
help                   | Prints this help
		`)
	}
	var platformParam = flag.StringP("platform", "p", runtime.GOOS, "sets the target platform of the operation")
	flag.Parse()
	var command = flag.Arg(0)
	switch command {
	case "init":
		var location = filepath.Clean(flag.Arg(1))
		InitProject(location, *platformParam)
	case "watch":
		var location = filepath.Clean(flag.Arg(1))
		Watch(location)
	case "version":
		fmt.Printf("djs %s\n", version)
	case "bundle":
		var bundleName = filepath.Clean((flag.Arg(1)))
		Bundle(bundleName, *platformParam)

	case "help":
		fallthrough
	default:
		flag.Usage()
	}

	var showHelp bool
	flag.BoolVarP(&showHelp, "help", "h", false, "Prints this help")
	if showHelp {
		flag.Usage()
	}
}
