package main

import (
	"fmt"
	"os"
	"strings"
)

var version = "0.1.0"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println(version)
			os.Exit(0)
		case "-h", "--help":
			fmt.Println(help())
			os.Exit(0)
		}
	}

	bot := &Bot{Input: os.Stdin, Output: os.Stdout, ErrorOutput: os.Stderr}
	bot.Run()
}

func help() string {
	lines := []string{}
	lines = append(lines, "Usage:")
	lines = append(lines, "  shbot")
	lines = append(lines, "  shbot (--version | -v)")
	lines = append(lines, "  shbot (--help | -h)")
	lines = append(lines, "")
	lines = append(lines, "Options:")
	lines = append(lines, "  --version, -v\tShow version number")
	lines = append(lines, "  --help, -h\tShow help message")

	return strings.Join(lines, "\n")
}
