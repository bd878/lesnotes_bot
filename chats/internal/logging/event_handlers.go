package logging

import (
	"context"

	"go.uber.org/zap"

	"github.com/bd878/lesnotes_bot/internal/ddd"
	"github.com/bd878/lesnotes_bot/internal/logger"
)

type domainHandlers[T ddd.Event] struct {
	ddd.EventHandler[T]
	logger *logger.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*domainHandlers[ddd.Event])(nil)

func LogDomainEventHandlers[T ddd.Event](handler ddd.EventHandler[T], logger *logger.Logger) *domainHandlers[T] {
	return &domainHandlers[T]{
		EventHandler: handler,
		logger: logger,
	}
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	h.logger.Infof("--> Chats.On(%s)\n", event.EventName())
	defer func() { h.logger.WithOptions(zap.Fields(zap.Error(err))).Infof("<-- Chats.On(%s)\n", event.EventName()) }()
	return h.EventHandler.HandleEvent(ctx, event)
}