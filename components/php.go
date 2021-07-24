package components

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	archPHP "scythe.fr/cli-scythe/template"
)

func CreatePhpProject() {

	var rootDir string

	color.Blue("What is the name of your project? : ")
	fmt.Scanf("%s", &rootDir)

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
