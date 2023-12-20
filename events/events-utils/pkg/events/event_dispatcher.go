package events

import (
	"errors"
	"slices"
	"sync"
)

var (
	ErrEventAlreadyRegistered   = errors.New("handler already registered to this event")
	ErrEmptyListOfEventsToClean = errors.New("the list of events registered is empty")
	ErrNotRegisteredEvent       = errors.New("the event dispatcher was not register")
	ErrNoneHandlersToEvent      = errors.New("the event dispatcher has not handlers associated")
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

func (evd *EventDispatcher) Dispatch(event IEvent) error {
	eventHandlers := evd.handlers[event.GetName()]

	if eventHandlers == nil {
		return ErrNotRegisteredEvent

	}

	if len(eventHandlers) == 0 {
		return ErrNoneHandlersToEvent
	}

	wg := sync.WaitGroup{}

	wg.Add(len(eventHandlers))

	for _, handler := range eventHandlers {

		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			handler.Handle(event)
		}(&wg)

	}

	go func() { wg.Wait() }()

	return nil
}
