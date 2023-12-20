package events

import (
	"errors"
	"slices"
)

var (
	ErrEventAlreadyRegistered   = errors.New("handler already registered to this event")
	ErrEmptyListOfEventsToClean = errors.New("the list of events registered is empty")
)

type EventDispatcher struct {
	handlers map[string][]IEventHandler
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]IEventHandler),
	}
}

func (evd *EventDispatcher) Register(name string, handler IEventHandler) error {
	if _, ok := evd.handlers[name]; ok {
		for _, h := range evd.handlers[name] {
			if h == handler {
				return ErrEventAlreadyRegistered
			}
		}
	}

	evd.handlers[name] = append(evd.handlers[name], handler)

	return nil
}

func (evd *EventDispatcher) Clear() error {
	if len(evd.handlers) == 0 {
		return ErrEmptyListOfEventsToClean
	}

	evd.handlers = make(map[string][]IEventHandler)

	return nil
}

func (evd *EventDispatcher) Has(eventName string, handler IEventHandler) bool {
	handlers := evd.handlers[eventName]

	if handlers == nil {
		return false
	}

	return slices.Contains(handlers, handler)
}
