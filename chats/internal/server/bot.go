package server

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/go-telegram/bot/models"
	botApi "github.com/go-telegram/bot"

	"github.com/bd878/lesnotes_bot/internal/bot"
	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/internal/i18n"
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

	return nil
}

func (s server) CreateChat(ctx context.Context, b *botApi.Bot, update *models.Update) {
	err := s.app.CreateChat(ctx, application.CreateChat{
		Lang: i18n.LangRu,
		Login: fmt.Sprintf("%d", update.Message.Chat.ID),
		Password: uuid.New().String(),
		Chat: &update.Message.Chat,
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
