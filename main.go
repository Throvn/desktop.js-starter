package main

import (
	_ "embed"
	"fmt"

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
		var location = flag.Arg(2)
		InitProject(name, location)
	case "watch":
		var location = flag.Arg(1)
		Watch(location)
	case "version":
		fmt.Printf("djs %s\n", version)
	}
}
