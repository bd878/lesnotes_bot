package domain

import (
	"errors"

	"github.com/go-telegram/bot/models"
	"github.com/bd878/lesnotes_bot/internal/i18n"
)

var (
	ErrChatExists = errors.New("chat exists")
	ErrChatEmpty = errors.New("chat is nil")
)

type Chat struct {
	ID string
	Chat *models.Chat
	Lang i18n.LangCode
}

func NewChat(id string) *Chat {
	return &Chat{ID: id}
}

func CreateChat(id string, chat *models.Chat) (*Chat, error) {
	c := NewChat(id)
	if chat == nil {
		return nil, ErrChatEmpty
	}
	c.Chat = chat
	return c, nil
}