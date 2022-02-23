package cmd

// TODO : Add a persistent flag to init a repository
// ! Multiple colors on the select prompt

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This CLI is used to generate boilerplate code for a lot of projects easily.",
	Long:  `A CLI only made by Scythe and for Scythe to tidy up my work (Of course I am Scythe ! Contact me at Scythe@outlook.fr)`,
	Run: func(cmd *cobra.Command, args []string) {
		flagState, _ := cmd.Flags().GetBool("git")

		promptValidation(flagState)

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolP("git", "g", false, "Init a repository easily")

}

func promptValidation(flagState bool) {

	rootDir, err := os.Getwd()

	color.HiYellow("\n?" + "" + " Choose one type of project")

	if err != nil {
		fmt.Println(err)
	}

	color.Blue("\nWhat is the name of your project ? : ")
	fmt.Scanf("%s", &rootDir)

	if flagState {
		exec.Command("git", "init", rootDir)
	}

	exec.Command("code", rootDir).Start()

}
