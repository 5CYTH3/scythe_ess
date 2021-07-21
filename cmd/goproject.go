package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// goprojectCmd represents the goproject command
var goprojectCmd = &cobra.Command{
	Use:   "go",
	Short: "Allows you to create a full functional go project",
	Long:  `Everything is in the short maaaan`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goproject called")
	},
}

func init() {
	createCmd.AddCommand(goprojectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goprojectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	goprojectCmd.Flags().BoolP("complex", "c", true, "Set project type to complex")

}

func initGoProject() {

	var rootDir string

	color.Blue("\nWhat is the name of your project ? : ")
	fmt.Scanf("%s", &rootDir)

	if goprojectCmd.Flags().Changed("complex") {

	} else {

	}

}

func createGoProject(rootDir string) {
	err := os.MkdirAll(rootDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
