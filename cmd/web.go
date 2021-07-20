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

	color.Blue("What is the name of your project ? : ")
	fmt.Scanf("%s", &resp)

	// Create main root file
	err := os.MkdirAll(resp, os.ModePerm)

	// Create subdir
	os.MkdirAll(resp+"/css/", os.ModePerm)
	os.Create(resp + "css/style.css")
	os.MkdirAll(resp+"/res/images", os.ModePerm)
	os.MkdirAll(resp+"/res/video", os.ModePerm)

	// Create style.css file
	createStylesheetFile(resp + "/css/style.css")

	// Create index.html file
	createIndexFile(resp + "/index.html")

	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command("code", resp)
	cmd.Start()
}

func createSassBoilerplate() {
	fmt.Println("SASS BOILERPLATE FUNCTION CALLED")

}

func createStylesheetFile(path string) {

	file, erro := os.Create(path)
	if erro != nil {
		fmt.Println(erro)
	}
	file.WriteString(
		`* {
	padding: 0;
	margin: 0;
}
body {
	width: 100%;
	list-style-type: none;
	text-decoration: none;
}
`)

	defer file.Close()

}

func createIndexFile(path string) {
	file, erro := os.Create(path)

	if erro != nil {
		fmt.Println(erro)
	}

	file.WriteString(
		`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <main>
		<header class="header" id="header">
			<h1 id="hero__title">
			<ul id="nav">
				<li></li>
				<li></li>
				<li></li>
			</ul>
		</header>
		<section class="home__section" id="home">
		
		</section>
	</main>
</body>
</html>
`)

}
