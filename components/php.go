package components

import (
	"os"

	archPHP "scythe.fr/cli-scythe/template"
)

func CreatePhpProject(rootDir string) {

	err := os.Mkdir(rootDir, os.ModePerm)

	if err != nil {
		panic(err)
	}

	createPhpArchitecture(rootDir)

}

func createPhpArchitecture(rootPath string) {

	createDirs := []string{
		rootPath + "/sass/pages",
		rootPath + "/sass/components",
		rootPath + "/sass/models",
		rootPath + "/css",
		rootPath + "/js",
		rootPath + "/php/components",
		rootPath + "/php/models",
		rootPath + "/php/controllers",
		rootPath + "/php/pages",
	}

	for _, dir := range createDirs {

		os.MkdirAll(dir, os.ModePerm)

	}

	archPHP.CreateAwesomeArchitecture(rootPath + "/sass")
	archPHP.CreatePHPfiles(rootPath)

}
