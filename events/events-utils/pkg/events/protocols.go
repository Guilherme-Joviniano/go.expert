package events

import "time"

type IEvent interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type IEventHandler interface {
	Handle(event IEvent) error
}

type IEventDispatcher interface {
	Has(eventName string, handler IEventHandler) bool
	Remove(name string, handler IEventHandler) error
	Register(name string, handler IEventHandler) error
	Dispatch(event IEvent) error
	Clear() error
}
