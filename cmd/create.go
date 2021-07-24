package cmd

// ! Maybe put the prompt on create command to choose which project create with it. But flags will be useless !?
// TODO : Add a persistent flag to init a repository

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	projects "scythe.fr/cli-scythe/components"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This CLI is used to generate boilerplate code for a lot of projects easily.",
	Long:  `A CLI only made by Scythe and for Scythe to tidy up my work (Of course I am Scythe ! Contact me at Scythe@outlook.fr)`,
	Run: func(cmd *cobra.Command, args []string) {

		promptValidation()

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}

func showPrompt() string {
	items := []string{
		"CSS Web Boilerplate",
		"SASS Web Boilerplate",
		"PHP + SASS Web Boilerplate",
		"Discord.js pre-built project",
		"Golang pre-built project",
	}
	index := -1

	var result string

	template := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U000027A4 {{ .| white}}",
		Inactive: "  {{ . | red | blue | magenta | yellow | cyan }}",
		Help:     "[Use arrow keys]",
	}

	prompt := promptui.Select{
		Label:     "\n?" + "" + " Choose one type of project",
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

func promptValidation() {

	typeProject := showPrompt()
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	var rootDir string
	color.Blue("\nWhat is the name of your project ? : ")
	fmt.Scanf("%s", &rootDir)

	if typeProject == "CSS Web Boilerplate" {
		projects.CreateCssBoilerplate(rootDir)
		color.Green("\nYour CSS project has been generated in " + path)
	} else if typeProject == "SASS Web Boilerplate" {
		projects.CreateSassBoilerplate(rootDir)
		color.Green("\nYour SASS project has been generated in " + path)
	} else if typeProject == "PHP + SASS Web Boilerplate" {
		projects.CreatePhpProject(rootDir)
	} else if typeProject == "Discord.js pre-built project" {
		projects.CreateDjsProject(rootDir)
	} else if typeProject == "Golang pre-built project" {
		projects.CreateGoProject(rootDir)
	}

	exec.Command("code", rootDir).Start()

}
