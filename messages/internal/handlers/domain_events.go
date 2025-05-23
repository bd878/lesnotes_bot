package handlers

import (
	"context"
	"fmt"

	botApi "github.com/go-telegram/bot"

	"github.com/bd878/lesnotes_bot/internal/bot"
	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/internal/ddd"
	"github.com/bd878/lesnotes_bot/messages/internal/domain"
)

type domainHandlers struct {
	bot *bot.Bot
	logger *logger.Logger
}

func NewDomainHandlers(bot *bot.Bot, logger *logger.Logger) *domainHandlers {
	return &domainHandlers{bot: bot, logger: logger}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handler ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handler, domain.MessageCreatedEvent)
}

func (h domainHandlers) HandleEvent(ctx context.Context, event ddd.Event) error {
	switch event.EventName() {
	case domain.MessageCreatedEvent:
		return h.onMessageCreatedEvent(ctx, event)
	}
	return nil
}

func (h domainHandlers) onMessageCreatedEvent(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.MessageCreated)

	_, err := h.bot.SendMessage(ctx, &botApi.SendMessageParams{
		ChatID: payload.Message.ChatID,
		Text: fmt.Sprintf("https://stage.lesnotes.space/m/%d", payload.Message.Message.ID),
	})
	if err != nil {
		h.logger.Errorln(err)
		return err
	}

	return nil
}