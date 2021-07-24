package cmd

// TODO : Add PHP projects (with SASS)
// ! Maybe put the prompt on create command to choose which project create with it. But flags will be useless !?
// TODO : Add a persistent flag to init a repository

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This CLI is used to generate boilerplate code for a lot of projects easily.",
	Long:  `A CLI only made by Scythe and for Scythe to tidy up my work (Of course I am Scythe ! Contact me at Scythe@outlook.fr)`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("\nUsage : command <commands> [@argument]")
		fmt.Println("Here are the types of projects you can create :")
		fmt.Println(" - Web Boilerplate (<command> = web)\n - Golang project (<command> = go)\n")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}
