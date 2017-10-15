package main

import (
	"os/exec"
	"strings"

	"github.com/naoty/shbot/adapters"
)

// Bot represents a bot interacting with users.
type Bot struct {
	Adapter adapters.Adapter
}

// Run starts to accept messages from bot.Adapter.
func (bot *Bot) Run() {
	for {
		bot.Adapter.Prepare()

		message := bot.Adapter.ReadMessage()
		if message == "" {
			continue
		}

		words := strings.Split(message, " ")
		name, args := words[0], words[1:]
		path, err := exec.LookPath(name)
		if err != nil {
			bot.Adapter.WriteError(err)
			continue
		}

		command := exec.Command(path, args...)
		out, err := command.Output()
		if err != nil {
			bot.Adapter.WriteError(err)
		}

		bot.Adapter.WriteMessage(string(out))
	}
}
