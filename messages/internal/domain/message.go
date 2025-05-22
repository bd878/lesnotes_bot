package domain

import (
	"errors"
	galleryMessages "github.com/bd878/gallery/server/messages/pkg/model"
)

var (
	ErrTextEmpty = errors.New("text empty")
	ErrUserIDEmpty = errors.New("user id empty") 
)

type Message struct {
	ID string
	Message *galleryMessages.Message
}

func NewMessage(id string) *Message {
	return &Message{ID: id}
}

func CreateMessage(id string, text string, userID int32) (*Message, error) {
	if text == "" {
		return nil, ErrTextEmpty
	}

	if userID == 0 {
		return nil, ErrUserIDEmpty
	}

	msg := NewMessage(id)
	msg.Message = &galleryMessages.Message{
		Text: text,
		UserID: userID,
	}

	return msg, nil
}