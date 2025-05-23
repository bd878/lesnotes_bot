package application

import (
	"context"

	"github.com/go-telegram/bot/models"

	"github.com/bd878/lesnotes_bot/chats/internal/domain"
	"github.com/bd878/lesnotes_bot/internal/ddd"
	"github.com/bd878/lesnotes_bot/internal/i18n"
)

type (
	CreateChat struct {
		Login string
		Password string
		Lang i18n.LangCode
		Chat *models.Chat
	}

	KickMember struct {
		ID int64
	}

	GetChat struct {
		ID int64
	}

	App interface {
		CreateChat(ctx context.Context, cmd CreateChat) error
		KickMember(ctx context.Context, cmd KickMember) error
		GetChat(ctx context.Context, query GetChat) (*domain.Chat, error)
	}

	Application struct {
		chats domain.ChatRepository
		gateway domain.ChatGateway
		publisher ddd.EventPublisher[ddd.Event]
	}
)

var _ App = (*Application)(nil)

func New(chats domain.ChatRepository, gateway domain.ChatGateway, publisher ddd.EventPublisher[ddd.Event]) *Application {
	return &Application{
		chats: chats,
		gateway: gateway,
		publisher: publisher,
	}
}

func (a Application) CreateChat(ctx context.Context, cmd CreateChat) error {
	token, err := a.gateway.Signup(ctx, cmd.Login, cmd.Password)
	if err != nil {
		return err
	}

	chat, err := domain.CreateChat(token, cmd.Login, cmd.Password, cmd.Lang, cmd.Chat)
	if err != nil {
		return err
	}

	err = a.chats.Save(ctx, chat)
	if err != nil {
		return err
	}

	return a.publisher.Publish(ctx, chat.Events()...)
}

func (a Application) GetChat(ctx context.Context, query GetChat) (*domain.Chat, error) {
	var (
		token string
		err error
	)

	chat, err := a.chats.Load(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	err = a.gateway.Auth(ctx, chat.Token)
	switch err {
	case domain.ErrExpired:
		token, err = a.gateway.Login(ctx, chat.Login, chat.Password)
		if err != nil {
			return nil, err
		}
		err = a.chats.UpdateToken(ctx, chat.Chat.ID, token)
		if err == nil {
			return nil, err
		}
	case nil:
	default:
		return nil, err
	}

	chat.Token = token

	return chat, nil
}

func (a Application) KickMember(ctx context.Context, cmd KickMember) error {
	var token string

	chat, err := a.chats.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = a.gateway.Auth(ctx, chat.Token)
	switch err {
	case domain.ErrExpired:
		token, err = a.gateway.Login(ctx, chat.Login, chat.Password)
		if err != nil {
			return err
		}
	case nil:
		token = chat.Token
	default:
		return err
	}

	err = a.chats.Remove(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = chat.Delete(); err != nil {
		return err
	}

	err = a.gateway.Delete(ctx, chat.Login, chat.Password, token)
	if err != nil {
		return err
	}

	return a.publisher.Publish(ctx, chat.Events()...)
}