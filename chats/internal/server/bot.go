package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/go-telegram/bot/models"
	botApi "github.com/go-telegram/bot"
	galleryUsers "github.com/bd878/gallery/server/users/pkg/model"

	"github.com/bd878/lesnotes_bot/internal/bot"
	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/chats/internal/application"
)

type server struct {
	app application.App
	logger *logger.Logger
}

func RegisterBot(app application.App, bot *bot.Bot, logger *logger.Logger) error {
	s := &server{app: app, logger: logger}

	bot.RegisterHandler(botApi.HandlerTypeMessageText, "/start", botApi.MatchTypeExact, s.CreateChat)
	bot.RegisterHandlerMatchFunc(memberKickedMatch, s.KickMember)
	bot.RegisterHandlerMatchFunc(messageTextMatch, s.CreateMessage)

	return nil
}

func (s server) CreateChat(ctx context.Context, b *botApi.Bot, update *models.Update) {
	id := uuid.New().String()

	err := s.app.CreateChat(ctx, application.CreateChat{ID: id, Chat: &update.Message.Chat})
	if err != nil {
		s.logger.Errorln(err)
		return
	}

	// TODO: replace all side effects on handlers
	_, err = b.SendMessage(ctx, &botApi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "created",
	})
	if err != nil {
		s.logger.Errorln(err)
		return
	}
}

func (s server) KickMember(ctx context.Context, b *botApi.Bot, update *models.Update) {
	err := s.app.KickMember(ctx, application.KickMember{})
	if err != nil {
		s.logger.Errorln(err)
		return
	}
}

func (s server) CreateMessage(ctx context.Context, b *botApi.Bot, update *models.Update) {
	id := uuid.New().String()

	err := s.app.CreateMessage(ctx, application.CreateMessage{
		ID: id,
		UserID: galleryUsers.PublicUserID,
		Text: update.Message.Text,
	})
	if err != nil {
		s.logger.Errorln(err)
		return
	}
}

func (s server) ConfirmIssue(ctx context.Context, b *botApi.Bot, update *models.Update) {
	err := s.app.ConfirmIssue(ctx, application.ConfirmIssue{})
	if err != nil {
		s.logger.Errorln(err)
		return
	}
}