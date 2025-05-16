package domain

import (
	"fmt"

	"github.com/go-telegram/bot/models"
	"github.com/bd878/lesnotes_bot/internal/i18n"
)

var (
	ErrChatExists = fmt.Errorf("chat exists")
)

type Chat struct {
	*models.Chat
	Lang i18n.LangCode
}

func NewChat() *Chat {
	return &Chat{Lang: i18n.LangRu}
}

func CreateChat(chat *models.Chat) (*Chat, error) {
	return &Chat{Chat: chat}, nil
}