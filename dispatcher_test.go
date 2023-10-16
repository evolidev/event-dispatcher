package event_dispatcher

import (
	"errors"
	"sync"
	"testing"
)

func TestDispatcherShouldCallAllListeners(t *testing.T) {
	callStack := make([]string, 0)
	p := &testProvider{listeners: make([]Listener, 0)}
	p.Add(ListenerFunc(func(event Event) (Event, error) {
		callStack = append(callStack, "func1")
		return event, nil
	}))
	p.Add(ListenerFunc(func(event Event) (Event, error) {
		callStack = append(callStack, "func2")
		return event, nil
	}))

	d := NewEventDispatcher(p)
	err := d.Dispatch(&testEvent{})

	if err != nil {
		t.Errorf("No error expected. Got: %s", err)
	}

	if len(callStack) != 2 {
		t.Errorf("Expected count %d. Got %d", 2, len(callStack))
	}
}

func TestDispatcherShouldCallListenersBeforeOneListenerStops(t *testing.T) {
	callStack := make([]string, 0)
	p := &testProvider{listeners: make([]Listener, 0)}
	p.Add(ListenerFunc(func(event Event) (Event, error) {
		callStack = append(callStack, "func1")
		return event.(StoppableEvent).WithPropagationStopped(), nil
	}))
	p.Add(ListenerFunc(func(event Event) (Event, error) {
		callStack = append(callStack, "func2")
		return event, nil
	}))

	d := NewEventDispatcher(p)
	err := d.Dispatch(&testStoppableEvent{stopped: false})

	if err != nil {
		t.Errorf("No error expected. Got: %s", err)
	}

	if len(callStack) != 1 {
		t.Errorf("Expected count %d. Got %d", 1, len(callStack))
	}
}

func TestDispatcherShouldReturnErrorIfHandlerReturnsOne(t *testing.T) {
	p := &testProvider{listeners: make([]Listener, 0)}
	p.Add(ListenerFunc(func(event Event) (Event, error) {
		return nil, errors.New("just a test")
	}))

	d := NewEventDispatcher(p)
	err := d.Dispatch(&testEvent{})

	if err == nil {
		t.Errorf("Expected error. Got nil")
	}
}

func TestGoHandlerShouldCallHandler(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	called := false

	p := &testProvider{listeners: make([]Listener, 0)}
	p.Add(ListenerFunc(func(event Event) (Event, error) {
		called = true
		wg.Done()
		return event, nil
	}))

	d := NewGoDispatcher(p)
	err := d.Dispatch(&testEvent{})

	if err != nil {
		t.Errorf("No error expected")
	}

	wg.Wait()
	if !called {
		t.Errorf("Handler got not called")
	}
}

type testStoppableEvent struct {
	stopped bool
}

func (t *testStoppableEvent) IsPropagationStopped() bool {
	return t.stopped
}

func (t *testStoppableEvent) WithPropagationStopped() StoppableEvent {
	return &testStoppableEvent{
		stopped: true,
	}
}

type testEvent struct {
}

type testProvider struct {
	listeners []Listener
}

func (t *testProvider) Add(listener Listener) {
	t.listeners = append(t.listeners, listener)
}

func (t *testProvider) GetListenersForEvent(event Event) []Listener {
	return t.listeners
}
