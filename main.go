package main

import (
	"fmt"
	"log"
	"os"

	"github.com/evanw/esbuild/pkg/api"
	flag "github.com/spf13/pflag"
)

func main() {
	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{"./tests/simple-jsx/source/index.tsx"},
		Outfile:           "./.internals/out/output.js",
		Bundle:            true,
		Write:             true,
		JSXFragment:       "\"group\"",
		JSXImportSource:   "libgui.so",
		Platform:          api.PlatformNeutral,
		Target:            api.ES2023,
		TreeShaking:       api.TreeShakingTrue,
		LogLevel:          api.LogLevelInfo,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Sourcemap:         api.SourceMapExternal,
		Metafile:          true,
		Color:             api.ColorIfTerminal,
		Tsconfig:          "./tsconfig.json",
		Supported: map[string]bool{
			"color-functions":          false,
			"gradient-double-position": false,
			"gradient-interpolation":   false,
			"gradient-midpoints":       false,
			"hwb":                      false,
			"hex-rgba":                 false,
			"inline-style":             false,
			"inset-property":           false,
			"is-pseudo-class":          false,
			"modern-rgb-hsl":           false,
			"nesting":                  false,
			"rebecca-purple":           false,
		},
	})
	fmt.Printf("%s", api.AnalyzeMetafile(result.Metafile, api.AnalyzeMetafileOptions{Color: true}))

	if len(result.Errors) > 0 {
		// os.Exit(1)
	}

	var nFlag = flag.String("n", "default", "help message for flag n")
	flag.Parse()
	var command = flag.Arg(0)
	if command == "init" {
		var name = flag.Arg(1)
		if name == "" {
			var err = fmt.Errorf("no project name given. Try specifying one: 'djs init <project name>'")
			log.Fatalf("%v", err)
			os.Exit(1)
		}

		var location = flag.Arg(2)
		InitProject(name, location)

	}

	fmt.Printf("Command: %s\n", command)
	fmt.Printf("Flag: %s\n", *nFlag)
}
