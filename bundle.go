package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func createAppBundleDarwin(name string) {

	var location, err = os.Getwd()
	var bundleLocation = filepath.Join(location, name+".app")
	if err != nil {
		Fatalf("Could not get working directory")
	}

	err = os.MkdirAll(filepath.Join(bundleLocation, "Contents", "MacOS"), 0755)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		Fatalf("Could not create appbundle at '%s'", location)
	}
	bundleLocation = filepath.Join(bundleLocation, "Contents")

	var enginePath = setupEngine(filepath.Join(bundleLocation, "MacOS"))
	fmt.Println("setup engine", enginePath)
	os.Rename(enginePath, filepath.Join(bundleLocation, "MacOS", name))

	err = os.MkdirAll(filepath.Join(bundleLocation, "Resources"), 0755)
	if err != nil {
		Fatalf("Could not create Resources folder")
	}

	pListFile, err := os.Create(filepath.Join(bundleLocation, "Info.plist"))
	if err != nil {
		Fatal("Could not create 'Info.plist' inside of appbundle which is required")
	}
	defer pListFile.Close()
	pListFile.WriteString(`
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>CFBundleName</key>
		<string>` + name + `</string>
		<key>CFBundleIdentifier</key>
		<string>org.js.desktop.` + name + `</string>
		<key>CFBundleVersion</key>
		<string>1.0</string>
		<key>CFBundleShortVersionString</key>
		<string>1.0</string>
		<key>CFBundleExecutable</key>
		<string>` + name + `</string>
	</dict>
</plist>
	`)

	sourceCodeLocation := filepath.Join(location, ".internals", "javascript", "index.js")
	sourceCode, err := os.ReadFile(sourceCodeLocation)
	if err != nil {
		Fatalf("Could not read js file at '%s'", sourceCodeLocation)
	}
	err = os.WriteFile(filepath.Join(bundleLocation, "Resources", "index.js"), sourceCode, 0644)
	if err != nil {
		Fatalf("Could not write js source code to app file")
	}
}

func deleteAppBundleDarwin(name string) {
	dir, err := os.Getwd()
	if err != nil {
		Fatalf("Could not get working directory\n%v", err)
	}
	err = os.RemoveAll(filepath.Join(dir, name+".app"))
	if err != nil {
		Fatalf("Could not remove previous app bundle\nExiting early\n%v", err)
	}
}

func Bundle(name string) {
	if name == "" || name == "." {
		Fatal("Please supply an app name.\nE.g. djs bundle MyCoolApp")
	}

	if runtime.GOOS != "darwin" {
		Fatalf("Can only bundle on macOS for macOS right now.\n\tPlease contribute if you have the ability to bring desktop.js to more platforms.")
	}
	deleteAppBundleDarwin(name)
	createAppBundleDarwin(name)
}
