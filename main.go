package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
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

	for {
		fmt.Print("shbot> ")

		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		str = strings.TrimRight(str, "\n")
		if str == "" {
			continue
		}

		words := strings.Split(str, " ")
		name, args := words[0], words[1:]
		path, err := exec.LookPath(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		command := exec.Command(path, args...)
		command.Stdout = os.Stdout
		err = command.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
