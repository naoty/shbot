package main

import (
	"fmt"
	"os"
)

var version = "0.1.0"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println(version)
			os.Exit(0)
		}
	}

	bot := &Bot{Input: os.Stdin, Output: os.Stdout, ErrorOutput: os.Stderr}
	bot.Run()
}
