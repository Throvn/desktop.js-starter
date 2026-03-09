package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func addAppIconLinux(bundleLocation string, name string) {
	rawIconPath := filepath.Join(bundleLocation, "..", "assets", "icon.png")
	if _, err := os.Stat(rawIconPath); errors.Is(err, os.ErrNotExist) {
		return
	}

	rawIcon, err := os.Open(rawIconPath)
	if err != nil {
		Fatalf("Could not get app icon.\n%v", err)
	}
	defer rawIcon.Close()

	iconPath := filepath.Join(bundleLocation, filepath.Base(name)+".png")
	file, err := os.Create(iconPath)
	if err != nil {
		Fatalf("Could not create '%s'\n%v", iconPath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, rawIcon)
	if err != nil {
		Fatalf("Could not write icon to '%s': %v", iconPath, err)
	}
}

func createAppImageLinux(name string) {
	var location, err = os.Getwd()
	var bundleLocation = filepath.Join(location, name+".AppDir")
	if err != nil {
		Fatalf("Could not get working directory")
	}

	// Set up path for engine executable
	err = os.MkdirAll(filepath.Join(bundleLocation, "usr", "bin"), 0755)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		Fatalf("Could not create AppDir at '%s'", location)
	}

	var enginePath = setupEngine(filepath.Join(bundleLocation, "usr", "bin"), "linux")
	fmt.Println("setup engine", enginePath)
	os.Rename(enginePath, filepath.Join(bundleLocation, "usr", "bin", filepath.Base(name)))

	addAppIconLinux(bundleLocation, name)

	desktopFile, err := os.Create(filepath.Join(bundleLocation, filepath.Base(name)+".desktop"))
	if err != nil {
		Fatal("Could not create '" + filepath.Base(name) + ".desktop' inside of appbundle which is required")
	}
	defer desktopFile.Close()
	desktopFile.WriteString(`
		Name=` + filepath.Base(name) + `
		Exec=` + filepath.Base(name) + `
		Icon=` + filepath.Base(name) + `
		Type=Application
		Categories=Utility;
	`)

	appRunFile, err := os.Create(filepath.Join(bundleLocation, "AppRun"))
	if err != nil {
		Fatalf("Could not create AppRun file in %s %v", bundleLocation, err)
	}
	defer appRunFile.Close()
	appRunFile.WriteString("#!/bin/sh\nexec \"./usr/bin/" + filepath.Base(name) + " run ./usr/share/javascript/index.js\"")

	sourceCodeLocation := filepath.Join(bundleLocation, "..", ".internals", "javascript", "index.js")
	sourceCode, err := os.ReadFile(sourceCodeLocation)
	if err != nil {
		Fatalf("Could not read js file at '%s'", sourceCodeLocation)
	}

	err = os.MkdirAll(filepath.Join(bundleLocation, "usr", "share", "javascript"), 0755)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		Fatalf("Could not create AppDir at '%s'", location)
	}
	err = os.WriteFile(filepath.Join(bundleLocation, "usr", "share", "javascript", "index.js"), sourceCode, 0644)
	if err != nil {
		Fatalf("Could not write js source code to app file")
	}
}

func deleteAppDirLinux(name string) {
	dir, err := os.Getwd()
	if err != nil {
		Fatalf("Could not get working directory\n%v", err)
	}
	err = os.RemoveAll(filepath.Join(dir, name+".AppDir"))
	if err != nil {
		Fatalf("Could not remove previous app bundle\nExiting early\n%v", err)
	}
}

func BundleLinux(name string) {
	deleteAppDirLinux(name)
	createAppImageLinux(name)
}
