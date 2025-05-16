package server

import (
	botApi "github.com/go-telegram/bot"

	"github.com/bd878/lesnotes_bot/internal/bot"
	"github.com/bd878/lesnotes_bot/chats/internal/logging"
)

func RegisterBot(app logging.App, bot *bot.Bot) error {
	bot.RegisterHandler(botApi.HandlerTypeMessageText, "/start", botApi.MatchTypeExact, app.Start)

	return nil
}
