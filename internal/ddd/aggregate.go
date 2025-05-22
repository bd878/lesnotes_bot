package ddd

import (
	"github.com/google/uuid"
)

type (
	AggregateNamer interface {
		AggregateName() string
	}

	Eventer interface {
		AddEvent(name string, payload EventPayload)
		Events() []Event
	}

	Aggregate struct {
		id string
		name string
		events []Event
	}
)

var _ interface {
	Eventer
	AggregateNamer
} = (*Aggregate)(nil)

func NewAggregate(name string) Aggregate {
	return Aggregate{
		id: uuid.New().String(),
		name: name,
		events: make([]Event, 0),
	}
}

func (a Aggregate) ID() string {
	return a.id
}

func (a Aggregate) AggregateName() string {
	return a.name
}

func (a *Aggregate) AddEvent(name string, payload EventPayload) {
	a.events = append(a.events, newEvent(name, payload))
}

func (a Aggregate) Events() []Event {
	return a.events
}