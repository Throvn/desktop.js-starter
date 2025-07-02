package main

import (
	_ "embed"
	"fmt"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

//go:embed version
var version string

func main() {

	flag.Parse()
	var command = flag.Arg(0)
	switch command {
	case "init":
		var name = flag.Arg(1)
		var location = filepath.Clean(flag.Arg(2))
		InitProject(name, location)
	case "watch":
		var location = filepath.Clean(flag.Arg(1))
		Watch(location)
	case "version":
		fmt.Printf("djs %s\n", version)
	}
}
