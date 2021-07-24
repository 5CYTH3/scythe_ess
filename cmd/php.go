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

// phpCmd represents the php command
var phpCmd = &cobra.Command{
	Use:   "php",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		createPhpProject()
	},
}

func init() {
	createCmd.AddCommand(phpCmd)

}

func createPhpProject() {

	var rootDir string

	color.Blue("What is the name of your project? : ")
	fmt.Scanf("%s", &rootDir)

	err := os.Mkdir(rootDir, os.ModePerm)

	if err != nil {
		panic(err)
	}

	createPhpArchitecture(rootDir)

	exec.Command("code", rootDir).Start()

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

}
