package components

import (
	"fmt"
	"os"
	"os/exec"

	archGo "scythe.fr/cli-scythe/template"
)

func CreateGoProject(rootDir string) {

	createGoArchitecture(rootDir)
	cmd := exec.Command("go", "mod", "init")
	cmd.Start()

}

func createGoArchitecture(rootDir string) {
	err := os.MkdirAll(rootDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	dirList := []string{"/vendor", "/interfaces"}

	for _, createDirs := range dirList {
		os.MkdirAll(rootDir+createDirs, os.ModePerm)
	}

	archGo.GoProjectContent(rootDir)
}
