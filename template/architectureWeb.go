package template

import (
	"fmt"
	"os"
)

func CreateStylesheetFile(path string) {

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

func CreateIndexFile(path string) {
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

	defer file.Close()

}

func CreateAwesomeArchitecture(root string) {

	// Create main scss file
	file, erro := os.Create(root + "/style.scss")
	if erro != nil {
		fmt.Println(erro)
	}

	file.WriteString(`/* To compile SASS into CSS, please use this command :

	sass sass/style.scss:css/style.css --watch

*/

@import "components/_var.scss";
@import "components/_mixins.scss";

@import "pages/index.scss";

* {
	margin: 0;
	padding: 0;
}

body {
	width: 100%;
	list-style-type: none;
	text-decoration: none;
}
	`)

	defer file.Close()

	// Array of paths
	sassFiles := []string{"/components/_var.scss", "/components/_mixins.scss", "/pages/index.scss"}

	// Create others files
	for _, createdFiles := range sassFiles {
		os.Create(root + createdFiles)
	}

}
