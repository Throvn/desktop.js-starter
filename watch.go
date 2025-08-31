package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/evanw/esbuild/pkg/api"
	fsnotify "github.com/fsnotify/fsnotify"
)

func build(location string) {

	api.Build(api.BuildOptions{
		EntryPoints:       []string{filepath.Join(location, "source/index.tsx")},
		Outfile:           filepath.Join(location, ".internals/javascript/index.js"),
		Bundle:            true,
		Write:             true,
		JSXFragment:       "\"group\"",
		JSXImportSource:   "GUI",
		Platform:          api.PlatformNeutral,
		Target:            api.ES2023,
		TreeShaking:       api.TreeShakingTrue,
		LogLevel:          api.LogLevelWarning,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Sourcemap:         api.SourceMapExternal,
		Metafile:          true,
		Color:             api.ColorIfTerminal,
		Tsconfig:          filepath.Join(location, "tsconfig.json"),
		External:          []string{"GUI"},
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

	fmt.Println()
	var title = "--- Engine Output "
	PrintDivider(title)
}

func setupWatcher(location string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start first engine
	stop := make(chan struct{})
	go startEngine(stop, location)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					PrintDivider("")
					log.Printf("[Watch] Change detected: %s", event.Name)

					// rebuild js
					build(location)

					// stop old engine
					close(stop)

					// start new engine
					stop = make(chan struct{})
					go startEngine(stop, location)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add source directory to watcher
	err = watcher.Add(filepath.Join(location, "source"))
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever
	select {}
}

func startEngine(stopChan <-chan struct{}, projectLocation string) {
	location, err := filepath.Abs(projectLocation)
	if err != nil {
		Fatalf("Could not get absolute path of '%s'", projectLocation)
	}

	var binaryLocation string
	if runtime.GOOS == "darwin" {
		binaryLocation = filepath.Join(location, ".internals/djs-aarch64-macos")
	} else {
		Fatalf("No engine for target: %s", runtime.GOOS)
	}

	absBinaryLocation, err := filepath.Abs(binaryLocation)
	if err != nil {
		Fatalf("Could not get absolute path of engine\n%v", err)
	}

	cmd := exec.Command(absBinaryLocation, "run", filepath.Join(location, ".internals/javascript/index.js"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		Fatalf("Failed to start engine: %v", err)
	}

	go func() {
		<-stopChan // wait for stop signal
		if err := cmd.Process.Kill(); err != nil {
			log.Printf("[Watch] Failed to kill process: %v", err)
		}
	}()
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

	build(projectLocation)
	setupWatcher(projectLocation)
}
