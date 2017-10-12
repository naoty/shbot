package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// Bot represents a bot interacting with users.
type Bot struct {
	Input       io.Reader
	Output      io.Writer
	ErrorOutput io.Writer
}

// Run starts to accept messages from bot.Input and send messages to bot.Output.
func (bot *Bot) Run() {
	for {
		fmt.Fprint(bot.Output, "shbot> ")

		reader := bufio.NewReader(bot.Input)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(bot.ErrorOutput, err)
		}

		str = strings.TrimRight(str, "\n")
		if str == "" {
			continue
		}

		words := strings.Split(str, " ")
		name, args := words[0], words[1:]
		path, err := exec.LookPath(name)
		if err != nil {
			fmt.Fprintln(bot.ErrorOutput, err)
			continue
		}

		command := exec.Command(path, args...)
		command.Stdout = bot.Output
		err = command.Run()
		if err != nil {
			fmt.Fprintln(bot.ErrorOutput, err)
		}
	}
}
