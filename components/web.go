package components

// TODO : Possibility to init a repository with a boolean

import (
	"fmt"
	"os"

	archWeb "scythe.fr/cli-scythe/template"
)

func CreateCssBoilerplate(rootDir string) {

	// Create main root file
	err := os.MkdirAll(rootDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	pathList := []string{rootDir + "/css/", rootDir + "/res/images", rootDir + "/res/video"}

	// Create subdir
	for _, currentCreatedDir := range pathList {
		os.MkdirAll(currentCreatedDir, os.ModePerm)

	}

	// Create style.css file
	archWeb.CreateStylesheetFile(rootDir + "/css/style.css")

	// Create index.html file
	archWeb.CreateIndexFile(rootDir + "/index.html")

}

func CreateSassBoilerplate(rootDir string) {

	// Create main root file
	err := os.MkdirAll(rootDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	pathList := []string{rootDir + "/css/", rootDir + "/sass/models", rootDir + "/sass/components", rootDir + "/sass/pages", rootDir + "/res/images", rootDir + "/res/video", rootDir + "/js"}

	// Created all folders
	for _, createdDirs := range pathList {
		os.MkdirAll(createdDirs, os.ModePerm)
	}

	// Create index.html file
	archWeb.CreateIndexFile(rootDir + "/index.html")

	// Create SCSS Architecture
	archWeb.CreateAwesomeArchitecture(rootDir + "/sass")
}
