package template

import "os"

func GoProjectContent(path string) {
	file, _ := os.Create(path + "main.go")

	file.WriteString(`
		package main
		import (
			"fmt"
			"os"
		)

		func main() {
			fmt.Println("Hey :)")
		}
	`)
	file.Close()

}
