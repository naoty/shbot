package adapters

// Adapter is I/O for bots.
type Adapter interface {
	Prepare()
	ReadMessage() string
	WriteMessage(string)
	WriteError(error)
}
