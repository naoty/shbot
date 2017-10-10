package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("shbot> ")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	str = strings.TrimRight(str, "\n")
	fmt.Println(str)

	os.Exit(0)
}
