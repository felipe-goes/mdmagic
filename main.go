package main

import (
	"fmt"
	"mdmagic/scrapper"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println("Please provide the url")
	case 2:
		scrapper.StackExchange(os.Args[1])
	default:
		fmt.Println("Too many arguments")
	}
}
