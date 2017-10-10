package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("shbot> ")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(str)

	os.Exit(0)
}
