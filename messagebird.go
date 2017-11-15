package sachet

import (
	"errors"

	"github.com/messagebird/go-rest-api"
)

type MessageBirdConfig struct {
	AccessKey string `yaml:"access_key"`
	Gateway   int    `yaml:"gateway"`
}

type MessageBird struct {
	client *messagebird.Client
	params messagebird.MessageParams
}

func NewMessageBird(config MessageBirdConfig) *MessageBird {
	return &MessageBird{
		client: messagebird.New(config.AccessKey),
		params: messagebird.MessageParams{
			Gateway: config.Gateway,
		},
	}
}

func (mb *MessageBird) Send(message Message) error {
	var err error
	switch message.Type {
	case "text":
		_, err = mb.client.NewMessage(message.From, message.To, message.Text, &mb.params)
		break
	case "voice":
		params := messagebird.VoiceMessageParams{
			Originator: message.From,
			Language:   message.Language,
			Voice:      message.Voice,
			Repeat:     message.Repeat,
		}
		_, err = mb.client.NewVoiceMessage(message.To, message.Text, &params)
		break
	default:
		err = errors.New("Unknown message 'type': " + message.Type)
	}

	return err
}
