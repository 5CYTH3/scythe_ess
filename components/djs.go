package components

import (
	"os"
	"os/exec"
)

// .env

func CreateDjsProject(rootDir string) {

	err := os.Mkdir(rootDir, os.ModePerm)

	if err != nil {
		panic(err)
	}

	createDjsArchitecture(rootDir)

}

func createDjsArchitecture(rootPath string) {

	// Folders path
	createDirs := []string{
		rootPath + "/database",
		rootPath + "/src/assets",
		rootPath + "/src/commands",
		rootPath + "/src/helpers",
		rootPath + "/src/events",
		rootPath + "/src/utils",
	}

	// Folders generation
	for _, dir := range createDirs {
		os.MkdirAll(dir, os.ModePerm)
	}

	cmdInit := exec.Command("npm", "init", "--force", "--yes")
	cmdInit.Dir = rootPath
	cmdInit.Start()

	suffixes := []string{
		"discord.js",
		"fs",
	}

	for _, suffix := range suffixes {
		cmdDepInstall := exec.Command("npm", "i", suffix)
		cmdDepInstall.Dir = rootPath
		cmdDepInstall.Start()
	}

}
