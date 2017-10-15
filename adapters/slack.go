package adapters

// Slack is an adapter connected with Slack.
type Slack struct {
}

// Prepare runs preparation before accepting messages.
func (adapter *Slack) Prepare() {
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
