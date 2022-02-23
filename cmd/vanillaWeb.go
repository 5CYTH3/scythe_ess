package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var vanillaWebCmd = &cobra.Command{
	Use:   "vanilla-web",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		fmt.Scanln(&name)
		createDirs(name)
		runAllCommands(name)
	},
}

func init() {
	createCmd.AddCommand(vanillaWebCmd)
}

func createDirs(name string) {

	dirs := []string{
		"/scss/global",
		"/scss/components",
		"/scss/utils",
		"/build",
		"/res",
	}

	for _, dir := range dirs {
		os.MkdirAll(name+dir, os.ModePerm)
	}

}

func runAllCommands(path string) {
	os.Chdir(path)
	exec.Command("npm", "i", "cssnano", "gulp", "browser-sync", "gulp-html-minifier", "gulp-postcss", "gulp-sass", "sass").Start()
	exec.Command("npm", "i", "-D", "gulp-imagemin").Start()
	exec.Command("npm", "init", "-y").Start()
	exec.Command("git", "init").Start()
}
