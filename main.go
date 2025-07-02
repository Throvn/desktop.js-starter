package main

import (
	flag "github.com/spf13/pflag"
)

func main() {

	flag.Parse()
	var command = flag.Arg(0)
	switch command {
	case "init":
		var name = flag.Arg(1)
		var location = flag.Arg(2)
		InitProject(name, location)
	case "watch":
		Watch()
	}

}
