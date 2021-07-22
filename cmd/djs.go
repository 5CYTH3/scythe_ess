/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// djsCmd represents the djs command
var djsCmd = &cobra.Command{
	Use:   "djs",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		createDjsProject()
	},
}

func init() {
	createCmd.AddCommand(djsCmd)

}

// .env

func createDjsProject() {

	var rootDir string

	color.Blue("What is the name of your project? : ")
	fmt.Scanf("%s", &rootDir)

	err := os.Mkdir(rootDir, os.ModePerm)

	if err != nil {
		panic(err)
	}

	createDjsArchitecture(rootDir)

	exec.Command("code", rootDir).Start()

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
