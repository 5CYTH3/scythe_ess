package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	archPHP "scythe.fr/cli-scythe/template"
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

	archPHP.CreateAwesomeArchitecture(rootPath + "/sass")
	archPHP.CreatePHPfiles(rootPath)

}
