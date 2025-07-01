package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
)

func setupFolders(location string) {
	var paths = []string{".internals", "assets", "assets/fonts", "source"}
	for i := 0; i < len(paths); i++ {
		var mkDirErr = os.MkdirAll(location+"/"+paths[i], os.ModePerm)
		if mkDirErr != nil {
			log.Fatalf("could not create 'assets' directory at '%s'\n%v", location, mkDirErr)
			os.Exit(1)
		}
	}
}

//go:embed templates/types.d.ts
var tstypes string

func setupTsTypes(location string) {
	var newTypes, tsconfigErr = os.Create(location + "/.internals/types.d.ts")
	if tsconfigErr != nil {
		log.Fatalf("could not create 'types.d.ts' at '%s/.internals'\n%v\n", location, tsconfigErr)
		os.Exit(1)
	}

	newTypes.WriteString(tstypes)
	newTypes.Close()
}

//go:embed templates/tsconfig.json
var tsconfig string

func setupTsConfig(location string) {
	var newTsConfig, tsconfigErr = os.Create(location + "/tsconfig.json")
	if tsconfigErr != nil {
		log.Fatalf("could not create 'tsconfig.json' at '%s'\n%v\n", location, tsconfigErr)
		os.Exit(1)
	}
	newTsConfig.WriteString(tsconfig)
	newTsConfig.Close()
}

func InitProject(name string, location string) {
	var projectLocation string
	if location == "" {
		var locErr error
		projectLocation, locErr = os.Getwd()
		if locErr != nil {
			var err = fmt.Errorf("could not resolve project location. Try specifying one: 'djs init %s <project location>'", name)
			log.Fatalf("%v\n%v", err, locErr)
			os.Exit(1)
		}
	}

	fmt.Printf("Info: Creating '%s' at '%s'\n", name, projectLocation)

	var projectDir, dirErr = os.ReadDir(projectLocation)
	if dirErr != nil {
		log.Fatalf("could not read directory contents of '%s'\n%v", projectLocation, dirErr)
		os.Exit(1)
	}
	if len(projectDir) > 0 {
		projectLocation += "/" + name
	}

	setupFolders(projectLocation)
	setupTsTypes(projectLocation)
	setupTsConfig(projectLocation)

}
