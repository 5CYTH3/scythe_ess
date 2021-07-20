package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "create":
			create()
		default:
			color.Red("Wrong arg ! Specify a command in this list :")
			scytheDefault()
		}
	} else {
		scytheDefault()
	}
}

func create() {
	create := flag.NewFlagSet("create", flag.ExitOnError)
	color.Blue("Choose the type of project you wanna create ->")
	create.Parse(os.Args)
}

func scytheDefault() {
	fmt.Println(" ")
	fmt.Println("Usage : scythe <command> [<args>]")
	fmt.Println("Available commands :")
	fmt.Println("- create")

}
