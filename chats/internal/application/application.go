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
		Start(ctx context.Context, b *botApi.Bot, update *models.Update)
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

func (a *Application) Start(ctx context.Context, b *botApi.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &botApi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "pong",
	})
	if err != nil {
		a.logger.Errorw("failed to send start message", "error", err)
	}
}