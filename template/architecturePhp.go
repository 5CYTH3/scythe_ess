package template

import (
	"fmt"
	"os"
)

func CreatePHPfiles(rootPath string) {

	indexFile, erro := os.Create(rootPath + "/index.php")

	if erro != nil {
		fmt.Println(erro)
	}

	indexFile.WriteString(
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

	defer indexFile.Close()

}
