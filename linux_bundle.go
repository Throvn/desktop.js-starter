package main

import (
	"fmt"
	"os"
	"path/filepath"
)

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
	os.Rename(enginePath, filepath.Join(bundleLocation, "usr", "bin", name))
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
