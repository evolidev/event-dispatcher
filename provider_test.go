package event_dispatcher

import "testing"

func TestTreeProviderShouldBeAbleToCallListenerForSpecificEvent(t *testing.T) {
	tlp := NewTreeProvider()
	tlp.Add("test.event.one", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))
	tlp.Add("test.event.two", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))

	l := tlp.GetListenersForEvent(&testEventOne{})

	if len(l) != 1 {
		t.Errorf("Expected exatly one listener")
	}
}

func TestTreeProviderShouldBeAbleToReturnAllListenerForSubType(t *testing.T) {
	tlp := NewTreeProvider()
	tlp.Add("", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))
	tlp.Add("test", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))
	tlp.Add("test.event", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))
	tlp.Add("test.event.one", ListenerFunc(func(event Event) (Event, error) {
		return event, nil
	}))

	l := tlp.GetListenersForEvent(&testEventOne{})

	if len(l) != 4 {
		t.Errorf("Expected exatly %d listener", 4)
	}
}

type testEventOne struct {
}
