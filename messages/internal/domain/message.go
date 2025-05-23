package domain

import (
	"errors"
	"github.com/bd878/lesnotes_bot/internal/ddd"
	galleryMessages "github.com/bd878/gallery/server/messages/pkg/model"
)

const MessageAggregate = "messages.Message"

var (
	ErrTextEmpty = errors.New("text empty")
	ErrUserIDEmpty = errors.New("user id empty") 
)

type Message struct {
	ddd.Aggregate
	Message *galleryMessages.Message
	ChatID int64
}

func NewMessage() *Message {
	return &Message{
		Aggregate: ddd.NewAggregate(MessageAggregate),
	}
}

func CreateMessage(id int32, galleryMessage *galleryMessages.Message, chatID int64) (*Message, error) {
	if galleryMessage.Text == "" {
		return nil, ErrTextEmpty
	}

	if galleryMessage.UserID == 0 {
		return nil, ErrUserIDEmpty
	}

	message := NewMessage()

	galleryMessage.ID = id
	message.Message = galleryMessage
	message.ChatID = chatID

	message.AddEvent(MessageCreatedEvent, &MessageCreated{
		Message: message,
	})

	return message, nil
}