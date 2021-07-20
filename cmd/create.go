package cmd

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
