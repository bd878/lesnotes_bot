package domain

import (
	"errors"

	"github.com/go-telegram/bot/models"
	"github.com/bd878/lesnotes_bot/internal/ddd"
	"github.com/bd878/lesnotes_bot/internal/i18n"
)

const ChatAggregate = "chats.Chat"

var (
	ErrChatExists = errors.New("chat exists")
	ErrChatEmpty = errors.New("chat is nil")
	ErrTokenEmpty = errors.New("token is nil")
	ErrLangEmpty = errors.New("lang is empty")
	ErrLoginEmpty = errors.New("login is empty")
	ErrPasswordEmpty = errors.New("password is empty")
)

type Chat struct {
	ddd.Aggregate
	Chat *models.Chat
	Token string
	Login string
	Password string
	Lang i18n.LangCode
}

func NewChat() *Chat {
	return &Chat{
		Aggregate: ddd.NewAggregate(ChatAggregate),
	}
}

func CreateChat(token, login, password string, lang i18n.LangCode, botChat *models.Chat) (*Chat, error) {
	if botChat == nil {
		return nil, ErrChatEmpty
	}
	if token == "" {
		return nil, ErrTokenEmpty
	}
	if string(lang) == "" {
		return nil, ErrLangEmpty
	}
	if login == "" {
		return nil, ErrLoginEmpty
	}
	if password == "" {
		return nil, ErrPasswordEmpty
	}

	chat := NewChat()
	chat.Chat = botChat
	chat.Token = token
	chat.Login = login
	chat.Password = password
	chat.Lang = lang

	chat.AddEvent(ChatCreatedEvent, &ChatCreated{
		Chat: chat,
	})

	return chat, nil
}

func (Chat) Key() string { return ChatAggregate }
