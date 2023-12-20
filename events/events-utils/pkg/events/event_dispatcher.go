package events

import "errors"

var (
	ErrEventAlreadyRegistered = errors.New("handler already registered to this event")
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
