package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]
	switch command {
	case "parse_list":
		parseList()
	case "serve":
		upload()
	case "migrate":
		upload()
	default:
		fmt.Println("Unknown command")
	}
}
