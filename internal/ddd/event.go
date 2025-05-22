package ddd

import (
	"time"
	"github.com/google/uuid"
)

type (
	EventPayload interface {}

	Event interface {
		IDer
		EventName() string
		Payload() EventPayload
		OccurredAt() time.Time
	}

	IDer interface {
		ID() string
	}

	event struct {
		id string
		name string
		payload EventPayload
		occurredAt time.Time
	}
)

var _ Event = (*event)(nil)

func NewEvent(name string, payload EventPayload) event {
	return newEvent(name, payload)
}

func newEvent(name string, payload EventPayload) event {
	return event{
		id: uuid.New().String(),
		name: name,
		payload: payload,
		occurredAt: time.Now(),
	}
}

func (e event) ID() string { return e.id }
func (e event) EventName() string { return e.name }
func (e event) Payload() EventPayload { return e.payload }
func (e event) OccurredAt() time.Time { return e.occurredAt }
