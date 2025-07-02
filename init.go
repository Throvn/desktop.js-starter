package main

import (
	_ "embed"
	"fmt"
	"os"
)

func setupFolders(location string) {
	var paths = []string{".internals", "assets", "assets/fonts", "source"}
	for i := 0; i < len(paths); i++ {
		var mkDirErr = os.MkdirAll(location+"/"+paths[i], os.ModePerm)
		if mkDirErr != nil {
			Fatalf("Could not create 'assets' directory at '%s'\n%v", location, mkDirErr)
		}
	}
}

//go:embed templates/types.d.ts
var tstypes string

//go:embed templates/libgui.d.mts
var tsguitypes string

func setupTsTypes(location string) {
	var newTypes, tstypeErr = os.Create(location + "/.internals/types.d.ts")
	if tstypeErr != nil {
		Fatalf("Could not create 'types.d.ts' at '%s/.internals'\n%v\n", location, tstypeErr)
	}

	newTypes.WriteString(tstypes)
	newTypes.Close()

	var newGuiTypes, tsguitypeErr = os.Create(location + "/.internals/libgui.d.mts")
	if tsguitypeErr != nil {
		Fatalf("Could not create 'libgui.d.mts' at '%s/.internals'\n%v\n", location, tsguitypeErr)
	}

	newGuiTypes.WriteString(tsguitypes)
	newGuiTypes.Close()
}

//go:embed templates/tsconfig.json
var tsconfig string

func setupTsConfig(location string) {
	var newTsConfig, tsconfigErr = os.Create(location + "/tsconfig.json")
	if tsconfigErr != nil {
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
		Fatalf("Could not create '/source/index.tsx' at '%s'\n%v\n", location, tsconfigErr)
	}
	newTsConfig.WriteString(tsindex)
	newTsConfig.Close()
}

func InitProject(name string, location string) {
	if name == "" {
		Fatalf("no project name given. Try specifying one: \"djs init <project name>\"")
	}

	var projectLocation string = location
	if projectLocation == "" {
		var locErr error
		projectLocation, locErr = os.Getwd()
		fmt.Println(projectLocation)
		if locErr != nil {
			var err = fmt.Errorf("could not resolve project location. Try specifying one: 'djs init %s <project location>'", name)
			Fatalf("%v\n%v", err, locErr)
		}
	}

	if _, err := os.Stat(projectLocation); !os.IsNotExist(err) {
		Infof("Nothing done since project '%s' at '%s' already exists", name, projectLocation)
		os.Exit(0)
	}

	fmt.Printf("Info: Creating '%s' at '%s'\n", name, projectLocation)

	var projectDir, dirErr = os.ReadDir(projectLocation)
	if dirErr != nil {
		Fatalf("could not read directory contents of '%s'\n%v", projectLocation, dirErr)
	}
	if len(projectDir) > 0 {
		projectLocation += "/" + name
	}

	setupFolders(projectLocation)
	setupTsTypes(projectLocation)
	setupTsConfig(projectLocation)
	setupTsIndex(projectLocation)

}
