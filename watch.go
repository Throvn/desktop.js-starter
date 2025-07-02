package main

import (
	"fmt"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func Watch() {
	location, err := os.Getwd()
	if err != nil {
		Fatal("Could not get current working directory")
	}

	Infof("Starting to watch \"%s/source\" for changes\n", location)
	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{location + "/source/index.tsx"},
		Outfile:           "./.internals/out/output.js",
		Bundle:            true,
		Write:             true,
		JSXFragment:       "\"group\"",
		JSXImportSource:   "libgui",
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
		os.Exit(1)
	}
}
