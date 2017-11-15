package sachet

type Provider interface {
	Send(message Message) error
}

type Message struct {
	To       []string
	From     string
	Text     string
	Type     string
	Language string
	Voice    string
	Repeat   int
}

func NewMessage(to []string, from, text, msgType string) Message {
	message := Message{
		To:       to,
		From:     from,
		Text:     text,
		Type:     msgType,
		Language: "en-us",
		Voice:    "female",
		Repeat:   2,
	}
	return message
}
