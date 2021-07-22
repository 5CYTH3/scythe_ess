package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	archGo "scythe.fr/cli-scythe/template"
)

// goprojectCmd represents the goproject command
var goprojectCmd = &cobra.Command{
	Use:   "go",
	Short: "Allows you to create a full functional go project",
	Long:  `Everything is in the short maaaan`,
	Run: func(cmd *cobra.Command, args []string) {
		flagState, err := cmd.Flags().GetBool("complex")
		if err != err {
			panic(err)
		}
		initGoProject(flagState)
	},
}

func init() {
	createCmd.AddCommand(goprojectCmd)
	goprojectCmd.Flags().BoolP("complex", "c", false, "Set project type to complex")

}

func initGoProject(flagState bool) {

	var rootDir string

	color.Blue("\nWhat is the name of your project ? : ")
	fmt.Scanf("%s", &rootDir)

	if flagState == true {
		createGoComplexProject(rootDir)

	} else {
		createGoProject(rootDir)

	}
	cmd := exec.Command("go", "mod", "init")
	cmd.Start()
	cmdCode := exec.Command("code", rootDir)
	cmdCode.Start()

}

func createGoProject(rootDir string) {
	err := os.MkdirAll(rootDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	archGo.GoProjectContent(rootDir)

}

func createGoComplexProject(rootDir string) {
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
