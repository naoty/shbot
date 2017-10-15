package adapters

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

// Slack is an adapter connected with Slack.
type Slack struct {
	session *slack.RTM
}

// NewSlack returns a new Slack adapter.
func NewSlack() *Slack {
	token := os.Getenv("SLACK_ACCESS_TOKEN")
	client := slack.New(token)
	client.SetDebug(true)

	session := client.NewRTM()
	go session.ManageConnection()

	return &Slack{
		session: session,
	}
}

// Prepare runs preparation before accepting messages.
func (adapter *Slack) Prepare() {
}

// ReadMessage returns messages from input.
func (adapter *Slack) ReadMessage() string {
	select {
	case message := <-adapter.session.IncomingEvents:
		switch event := message.Data.(type) {
		case *slack.MessageEvent:
			return event.Msg.Text
		}
	}

	return ""
}

// WriteMessage send a given message to output.
func (adapter *Slack) WriteMessage(message string) {
	// TODO: Send messages to Slack.
	fmt.Println(message)
}

// WriteError send an error message to output.
func (adapter *Slack) WriteError(err error) {
	// TODO: Send error messages to Slack.
	fmt.Fprintln(os.Stderr, err)
}
