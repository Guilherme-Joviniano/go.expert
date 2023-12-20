package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Mock Dependencies //

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct{}

func (t *TestEventHandler) Handle(event IEvent) error {
	return nil
}

// Mock Dependencies //

type EventDispatcherTypeSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTypeSuite) SetupTest() {
	suite.event = TestEvent{Name: "any_name", Payload: "any_payload"}
	suite.event2 = TestEvent{Name: "any_name_2", Payload: "any_payload_2"}
	suite.handler = TestEventHandler{}
	suite.handler2 = TestEventHandler{}
	suite.handler3 = TestEventHandler{}
	suite.eventDispatcher = NewEventDispatcher()
}

func (s *EventDispatcherTypeSuite) TestEventDispatcher_Register() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	assert.Nil(s.T(), err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
}

func (s *EventDispatcherTypeSuite) TestEventDispatcher_RegisterCheckHandlerAssignment() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	assert.Nil(s.T(), err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
	assert.Equal(s.T(), s.eventDispatcher.handlers[s.event.GetName()][0], &s.handler)
}
func (s *EventDispatcherTypeSuite) TestEventDispatcher_RegisterWithSameHandlerEvent() {
	s.eventDispatcher.Register(s.event.GetName(), &s.handler)

	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)

	assert.Equal(s.T(), err, ErrEventAlreadyRegistered)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTypeSuite))
}
