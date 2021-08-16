package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "mv",
	Short: "Move a folder's content to another",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		destFlagState, _ := cmd.Flags().GetString("dest")
		stageFlagState, _ := cmd.Flags().GetString("stage")
		extFlagState, _ := cmd.Flags().GetString("ext")

		mainMove(destFlagState, stageFlagState, extFlagState)

	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
	moveCmd.Flags().String("dest", "", "Specify destination for the files")
	moveCmd.Flags().String("stage", "", "Specify folder which contains files")
	moveCmd.Flags().String("ext", "", "Add a specific extension filter")

}

func mainMove(destFlagState string, stageFlagState string, extFlagState string) {

	if destFlagState != "" {

		if stageFlagState != "" {

			moveFiles(destFlagState, stageFlagState, extFlagState)

		} else {
			fmt.Println("$ Please provide a valid initial dir.")
			os.Exit(1)
		}

	} else {

	}
}

func moveFiles(destination string, stage string, ext string) {

	dirs, _ := ioutil.ReadDir(stage)

	fmt.Println("$ Moved following files : ")

	if ext == "" {
		for _, dir := range dirs {

			err := os.Rename(stage+"/"+dir.Name(), destination+"/"+dir.Name())
			fmt.Println("- " + dir.Name())

			if err != nil {
				log.Fatal(err)
			}

		}
	} else {
		for _, dir := range dirs {
			if strings.HasSuffix(dir.Name(), ext) {

				err := os.Rename(stage+"/"+dir.Name(), destination+"/"+dir.Name())
				fmt.Println("- " + dir.Name())

				if err != nil {
					log.Fatal(err)
				}
			}

		}

	}

}
