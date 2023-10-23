package event_dispatcher

import "testing"

func TestTreeProviderShouldBeAbleToCallListenerForSpecificEvent(t *testing.T) {
	tlp := NewTreeProvider()
	tlp.Add("event_dispatcher.test.event.one", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))
	tlp.Add("event_dispatcher.test.event.two", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))

	l := tlp.GetListenersForEvent(&testEventOne{})

	if len(l) != 1 {
		t.Errorf("Expected exatly one listener. Got %d", len(l))
	}
}

func TestTreeProviderShouldBeAbleToReturnAllListenerForSubType(t *testing.T) {
	tlp := NewTreeProvider()
	tlp.Add("", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))
	tlp.Add("event_dispatcher.test", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))
	tlp.Add("event_dispatcher.test.event", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))
	tlp.Add("event_dispatcher.test.event.one", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))

	l := tlp.GetListenersForEvent(&testEventOne{})

	if len(l) != 4 {
		t.Errorf("Expected exatly %d listener. Got %d", 4, len(l))
	}
}

type testEventOne struct {
}
