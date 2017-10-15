package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/naoty/shbot/adapters"
)

var version = "0.1.0"

func main() {
	var adapter adapters.Adapter = adapters.NewShell()

	args := os.Args[1:]
	for i, arg := range args {
		switch arg {
		case "--version", "-v":
			fmt.Println(version)
			os.Exit(0)
		case "--help", "-h":
			fmt.Println(help())
			os.Exit(0)
		case "--adapter", "-a":
			if i < len(args)-1 {
				adapter = findAdapter(args[i+1])
			}
		}
	}

	bot := &Bot{Adapter: adapter}
	bot.Run()
}

func help() string {
	lines := []string{}
	lines = append(lines, "Usage:")
	lines = append(lines, "  shbot [--adapter | -a <adapter>]")
	lines = append(lines, "  shbot --version | -v")
	lines = append(lines, "  shbot --help | -h")
	lines = append(lines, "")
	lines = append(lines, "Options:")
	lines = append(lines, "  --adapter, -a <adapter> Specify adapter [default: shell]")
	lines = append(lines, "  --version, -v           Show version number")
	lines = append(lines, "  --help, -h              Show help message")

	return strings.Join(lines, "\n")
}

func findAdapter(name string) adapters.Adapter {
	switch name {
	case "shell":
		return adapters.NewShell()
	case "slack":
		return &adapters.Slack{}
	default:
		return adapters.NewShell()
	}
}
