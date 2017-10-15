package adapters

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

// Slack is an adapter connected with Slack.
type Slack struct {
	client *slack.Client
}

// NewSlack returns a new Slack adapter.
func NewSlack() *Slack {
	token := os.Getenv("SLACK_ACCESS_TOKEN")
	client := slack.New(token)
	client.SetDebug(true)

	return &Slack{
		client: client,
	}
}

// Prepare runs preparation before accepting messages.
func (adapter *Slack) Prepare() {
	_, err := adapter.client.GetGroups(false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("authenticated")
	os.Exit(0)
}

// ReadMessage returns messages from input.
func (adapter *Slack) ReadMessage() string {
	// TODO: Read messages from Slack.
	return ""
}

// WriteMessage send a given message to output.
func (adapter *Slack) WriteMessage(message string) {
	// TODO: Send messages to Slack.
}

// WriteError send an error message to output.
func (adapter *Slack) WriteError(err error) {
	// TODO: Send error messages to Slack.
}
