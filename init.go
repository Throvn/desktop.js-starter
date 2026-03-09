package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

func teardownProject(location string) {
	Warn("Reverting all file operations because of error")
	err := os.RemoveAll(location)
	if err != nil {
		Fatalf("Error while reverting: %v", err)
	}
}

func setupFolders(location string) {
	var paths = []string{".internals", ".internals/javascript", "assets", "assets/fonts", "source"}
	for i := 0; i < len(paths); i++ {
		var mkDirErr = os.MkdirAll(filepath.Join(location, paths[i]), 0o777)
		if mkDirErr != nil {
			teardownProject(location)
			Fatalf("Could not create 'assets' directory at '%s'\n%v", location, mkDirErr)
		}
	}
}

//go:embed templates/types.d.ts
var tstypes string

//go:embed templates/GUI.d.mts
var tsguitypes string

//go:embed templates/txiki.d.ts
var tstxikitypes string

func setupTsTypes(location string) {
	var newTypes, tstypeErr = os.Create(location + "/.internals/types.d.ts")
	if tstypeErr != nil {
		teardownProject(location)
		Fatalf("Could not create 'types.d.ts' at '%s/.internals'\n%v\n", location, tstypeErr)
	}

	newTypes.WriteString(tstypes)
	newTypes.Close()

	var newGuiTypes, tsguitypeErr = os.Create(location + "/.internals/GUI.d.mts")
	if tsguitypeErr != nil {
		teardownProject(location)
		Fatalf("Could not create 'GUI.d.mts' at '%s/.internals'\n%v\n", location, tsguitypeErr)
	}

	newGuiTypes.WriteString(tsguitypes)
	newGuiTypes.Close()

	var newTxikiTypes, tsTxikiTypeErr = os.Create(location + "/.internals/txiki.d.ts")
	if tsTxikiTypeErr != nil {
		teardownProject(location)
		Fatalf("Could not create 'txiki.d.mts' at '%s/.internals'\n%v\n", location, tsTxikiTypeErr)
	}

	newTxikiTypes.WriteString(tstxikitypes)
	newTxikiTypes.Close()
}

//go:embed templates/tsconfig.json
var tsconfig string

func setupTsConfig(location string) {
	var newTsConfig, tsconfigErr = os.Create(location + "/tsconfig.json")
	if tsconfigErr != nil {
		teardownProject(location)
		Fatalf("Could not create 'tsconfig.json' at '%s'\n%v\n", location, tsconfigErr)
	}
	newTsConfig.WriteString(tsconfig)
	newTsConfig.Close()
}

//go:embed templates/index.tsx
var tsindex string

func setupTsIndex(location string) {
	var newTsConfig, tsconfigErr = os.Create(location + "/source/index.tsx")
	if tsconfigErr != nil {
		teardownProject(location)
		Fatalf("Could not create '/source/index.tsx' at '%s'\n%v\n", location, tsconfigErr)
	}
	newTsConfig.WriteString(tsindex)
	newTsConfig.Close()
}

func setupGitignore(location string) {
	var file, err = os.Create(location + "/.gitignore")
	if err != nil {
		teardownProject(location)
		Fatalf("Could not create '/.gitignore' at '%s'\n%v\n", location, err)
	}
	file.WriteString(".internals/\n*.app\n")
	file.Close()
}

//go:embed templates/Roboto-Regular.ttf
var robotoFont []byte

func setupFonts(location string) {
	var file, err = os.Create(location + "/assets/fonts/Roboto-Regular.ttf")
	if err != nil {
		teardownProject(location)
		Fatalf("Could not create '/assets/fonts/Roboto-Regular.ttf' at '%s'\n%v\n", location, err)
	}
	file.Write(robotoFont)
	file.WriteString(`.internals/`)
	file.Close()
}

func setupIniFile(name string, location string) {
	var file, err = os.Create(location + "/djs.ini")
	if err != nil {
		teardownProject(location)
		Fatalf("Could not create '/djs.ini' at '%s'\n%v\n", location, err)
	}
	file.WriteString(`
[project]
name = ` + name + `
starter-version = ` + version + `

[window]
width = 600
height = 300
	`)
	file.Close()

}

//go:embed templates/djs-arm64-darwin
var macosEngine []byte

func setupEngine(location string) string {
	binPath := filepath.Join(location, "djs-arm64-darwin")
	file, err := os.Create(binPath)
	if err != nil {
		teardownProject(location)
		Fatalf("Could not create '%s'\n%v", binPath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(macosEngine))
	if err != nil {
		teardownProject(location)
		Fatalf("Could not write engine binary to '%s': %v", binPath, err)
	}

	err = file.Chmod(0o777)
	if err != nil {
		teardownProject(location)
		Fatalf("Could not make executable '%s': %v", binPath, err)
	}

	return binPath
}

func InitProject(location string) {
	absPath, absErr := filepath.Abs(location)
	if absErr != nil {
		Fatalf("Could not get absolute path of '%s'", location)
	}
	name := path.Base(absPath)

	// Determine location of the project root.
	if location == "" {
		var locErr error
		location, locErr = os.Getwd()
		fmt.Println(location)
		if locErr != nil {
			var err = fmt.Errorf("could not resolve project location. Try specifying one: 'djs init <location>'")
			Fatalf("%v\n%v", err, locErr)
		}
	}

	contents, err := os.ReadDir(location)
	if err != nil {
		Warnf("Could not read contents of directory '%s'", absPath)
		err = os.MkdirAll(absPath, 0o777)
		if err != nil {
			Fatalf("Creating project directory at '%s' failed", absPath)
		}
	}
	if len(contents) > 0 {
		Fatalf("Directory '%s' is not empty. Try specifying a location: 'djs init <location>'", name)
	}

	setupFolders(location)
	setupTsTypes(location)
	setupTsConfig(location)
	setupTsIndex(location)
	setupGitignore(location)
	setupIniFile(name, location)
	setupFonts(location)
	setupEngine(filepath.Join(location, ".internals"))

	Infof("Created '%s'", absPath)
	os.Chdir(location)
}
