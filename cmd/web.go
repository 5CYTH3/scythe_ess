package cmd

import (
	"fmt"
	"os"
	"os/exec"

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
		Label:     "?" + "" + " Choose one type of web boilerplate",
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

	if typeProject == "With CSS" {
		createCssBoilerplate()
	} else if typeProject == "With SASS" {
		createSassBoilerplate()
	}

}

func createCssBoilerplate() {

	var resp string

	fmt.Println("CSS BOILERPLATE FUNCTION CALLED")
	color.Blue("What is the name of your project ? : ")
	fmt.Scanf("%s", &resp)

	err := os.MkdirAll(resp+"/css/", os.ModePerm)
	os.MkdirAll(resp+"/res/images", os.ModePerm)
	os.MkdirAll(resp+"/res/video", os.ModePerm)

	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command("code", ".")
	cmd.Start()
}

func createSassBoilerplate() {
	fmt.Println("SASS BOILERPLATE FUNCTION CALLED")

}
