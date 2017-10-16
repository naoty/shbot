package adapters

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

// Slack is an adapter connected with Slack.
type Slack struct {
	client  *slack.Client
	rtm     *slack.RTM
	context *context
}

// NewSlack returns a new Slack adapter.
func NewSlack() *Slack {
	token := os.Getenv("SLACK_ACCESS_TOKEN")
	client := slack.New(token)

	rtm := client.NewRTM()
	go rtm.ManageConnection()

	return &Slack{
		client:  client,
		rtm:     rtm,
		context: nil,
	}
}

// context is a context which is alive between accepting and sending messages.
type context struct {
	channel string
}

// Prepare runs preparation before accepting messages.
func (adapter *Slack) Prepare() {
	adapter.context = &context{channel: ""}
}

// ReadMessage returns messages from input.
func (adapter *Slack) ReadMessage() string {
	select {
	case message := <-adapter.rtm.IncomingEvents:
		switch event := message.Data.(type) {
		case *slack.MessageEvent:
			adapter.context.channel = event.Channel
			return event.Msg.Text
		}
	}

	return ""
}

// WriteMessage send a given message to output.
func (adapter *Slack) WriteMessage(text string) {
	params := slack.NewPostMessageParameters()
	params.AsUser = true

	channel := adapter.context.channel
	adapter.client.PostMessage(channel, text, params)
}

// WriteError send an error message to output.
func (adapter *Slack) WriteError(err error) {
	// TODO: Send error messages to Slack.
	fmt.Fprintln(os.Stderr, err)
}
