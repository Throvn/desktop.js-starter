package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
	fsnotify "github.com/fsnotify/fsnotify"
)

func build(location string) {

	fmt.Println(strings.Repeat("-", 100) + "\n")

	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{filepath.Join(location, "index.tsx")},
		Outfile:           ".internals/out/output.js",
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
		Tsconfig:          "tsconfig.json",
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

	fmt.Printf("%s\n", api.AnalyzeMetafile(result.Metafile, api.AnalyzeMetafileOptions{Color: true}))

	fmt.Println(strings.Repeat("-", 100) + "\n")
	Info("Waiting for changes to rebuild...\n")
}

func setupWatcher(location string) {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					// log.Println("modified file:", event.Name)
					build(location)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add(location)
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

func Watch(location string) {
	var projectLocation string = location
	if projectLocation == "" {
		// TODO: Handle if location was given!
		var err error
		projectLocation, err = os.Getwd()
		if err != nil {
			Fatal("Could not get current working directory")
		}
	}

	_, fileName := filepath.Split(projectLocation)
	var sourceLocation = projectLocation
	if fileName != "source" {
		sourceLocation = filepath.Join(projectLocation, "source")
	}
	_, err := os.ReadDir(projectLocation)
	if err != nil {
		Fatalf("No directory \"%s\" found to watch.\nMaybe try \"djs watch <path>\"\n%v", sourceLocation, err)
	}

	fmt.Println()
	Infof("Starting to watch \"%s\" for changes...\n", sourceLocation)

	build(sourceLocation)
	setupWatcher(sourceLocation)
}
