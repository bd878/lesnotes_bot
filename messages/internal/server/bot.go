package server

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/go-telegram/bot/models"
	botApi "github.com/go-telegram/bot"
	galleryUsers "github.com/bd878/gallery/server/users/pkg/model"

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
	id := uuid.New().String()

	res, err := s.app.CreateMessage(ctx, application.CreateMessage{
		ID: id,
		UserID: galleryUsers.PublicUserID,
		Text: update.Message.Text,
	})
	if err != nil {
		s.logger.Errorln(err)
		return
	}

	_, err = b.SendMessage(ctx, &botApi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: fmt.Sprintf("https://stage.lesnotes.space/m/%d", res),
	})
	if err != nil {
		s.logger.Errorln(err)
		return
	}
}

func messageTextMatch(update *models.Update) bool {
	if update.Message != nil {
		return update.Message.Text != ""
	}
	return false
}
