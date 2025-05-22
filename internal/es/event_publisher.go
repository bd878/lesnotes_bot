package es

import (
	"context"
	"github.com/bd878/lesnotes_bot/internal/ddd"
)

type EventPublisher struct {
	publisher ddd.EventPublisher[ddd.Event]
}

func NewEventPublisher(publisher ddd.EventPublisher[ddd.Event]) EventPublisher {
	return EventPublisher{
		publisher: publisher,
	}
}

func (p EventPublisher) Save(ctx context.Context, event ddd.Event) error {
	return p.publisher.Publish(ctx, event)
}