package main

import (
	"fmt"
	"mdmagic/scrapper"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1, 2:
		fmt.Println("There are missing arguments")
	case 3:
		site := os.Args[1]
		url := os.Args[2]

		if site == "stackexchange" {
			scrapper.StackExchange(url)
		} else if site == "cppbyexample" {
			scrapper.CppByExample(url)
		}
	default:
		fmt.Println("Too many arguments")
	}
}
