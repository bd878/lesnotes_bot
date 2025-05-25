package handlers

import (
	"context"

	botApi "github.com/go-telegram/bot"

	"github.com/bd878/lesnotes_bot/internal/bot"
	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/internal/ddd"
	"github.com/bd878/lesnotes_bot/chats/internal/domain"
)

type domainHandlers struct {
	bot *bot.Bot
	logger *logger.Logger
}

func NewDomainHandlers(bot *bot.Bot, logger *logger.Logger) *domainHandlers {
	return &domainHandlers{bot: bot, logger: logger}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handler ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handler, domain.ChatCreatedEvent)
}

func (h domainHandlers) HandleEvent(ctx context.Context, event ddd.Event) error {
	switch event.EventName() {
	case domain.ChatCreatedEvent:
		return h.onChatCreatedEvent(ctx, event)
	case domain.ChatDeletedEvent:
		return h.onChatDeletedEvent(ctx, event)
	}
	return nil
}

func (h domainHandlers) onChatCreatedEvent(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.ChatCreated)

	_, err := h.bot.SendMessage(ctx, &botApi.SendMessageParams{
		ChatID: payload.Chat.Chat.ID,
		Text: "created",
	})
	if err != nil {
		h.logger.Errorln(err)
		return err
	}
	return nil
}

func (h domainHandlers) onChatDeletedEvent(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.ChatDeleted)

	h.logger.Infow("chat deleted", "id", payload.Chat.Chat.ID)

	return nil
}