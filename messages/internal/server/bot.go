package server

import (
	"context"

	"github.com/go-telegram/bot/models"
	botApi "github.com/go-telegram/bot"
	galleryUsers "github.com/bd878/gallery/server/users/pkg/model"
	galleryMessages "github.com/bd878/gallery/server/messages/pkg/model"

	"github.com/bd878/lesnotes_bot/internal/bot"
	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/messages/internal/application"
)

type server struct {
	app application.App
	logger *logger.Logger
}

func RegisterBot(app application.App, bot *bot.Bot, logger *logger.Logger) error {
	s := &server{app: app, logger: logger}

	bot.RegisterHandlerMatchFunc(messageTextMatch, s.CreateMessage)

	return nil
}

func (s server) CreateMessage(ctx context.Context, b *botApi.Bot, update *models.Update) {
	err := s.app.CreateMessage(ctx, application.CreateMessage{
		Message: &galleryMessages.Message{
			Text: update.Message.Text,
			UserID: galleryUsers.PublicUserID,
		},
		ChatID: update.Message.Chat.ID,
	})
	if err != nil {
		s.logger.Errorln(err)
	}
}

func messageTextMatch(update *models.Update) bool {
	if update.Message != nil {
		return update.Message.Text != ""
	}
	return false
}
