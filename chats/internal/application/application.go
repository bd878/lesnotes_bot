package application

import (
	"context"

	"github.com/go-telegram/bot/models"
	botApi "github.com/go-telegram/bot"

	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/chats/internal/domain"
)

type (
	App interface {
		Start(ctx context.Context, b *botApi.Bot, update *models.Update) error
	}

	Application struct {
		chats domain.ChatsRepository
		logger *logger.Logger
	}
)

var _ App = (*Application)(nil)

func New(chats domain.ChatsRepository, logger *logger.Logger) *Application {
	return &Application{
		chats: chats,
		logger: logger,
	}
}

func (a *Application) Start(ctx context.Context, b *botApi.Bot, update *models.Update) error {
	chat, err := domain.CreateChat(&update.Message.Chat)
	if err != nil {
		return err
	}

	err = a.chats.Save(ctx, chat)
	if err != nil {
		return err
	}

	_, err = b.SendMessage(ctx, &botApi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "created",
	})
	if err != nil {
		return err
	}
	return nil
}