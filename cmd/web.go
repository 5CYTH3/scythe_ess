package cmd

// TODO : Possibility to init a repository with a boolean

import (
	"fmt"
	"os"
	"os/exec"

	archWeb "scythe.fr/cli-scythe/template"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Allows you to make an html boilerplate with flags",
	Long:  `You can do HTML boilerplate folder with SASS or without, and with an architecture that I use frequently.`,
	Run: func(cmd *cobra.Command, args []string) {
		validateSelect()
	},
}

func init() {
	createCmd.AddCommand(webCmd)

}

func showSelectStyle() string {
	items := []string{"With SASS", "With CSS"}
	index := -1

	var result string

	template := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U000027A4 {{ .| cyan }}",
		Inactive: "  {{ . | blue }}",
		Help:     "[Use arrow keys]",
	}

	prompt := promptui.Select{
		Label:     "\n?" + "" + " Choose one type of web boilerplate",
		Items:     items,
		Templates: template,
	}

	index, result, err := prompt.Run()

	if index == -1 {
		items = append(items, result)
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return result
}

func validateSelect() {

	typeProject := showSelectStyle()
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	var rootDir string
	color.Blue("\nWhat is the name of your project ? : ")
	fmt.Scanf("%s", &rootDir)

	if typeProject == "With CSS" {
		createCssBoilerplate(rootDir)
		color.Green("\nYour CSS project has been generated in " + path)
	} else if typeProject == "With SASS" {
		createSassBoilerplate(rootDir)
		color.Green("\nYour SASS project has been generated in " + path)
	}

	cmd := exec.Command("code", rootDir)
	cmd.Start()

}

func createCssBoilerplate(rootDir string) {

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

func createSassBoilerplate(rootDir string) {

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
