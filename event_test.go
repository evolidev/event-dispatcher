package event_dispatcher

import "testing"

func TestName(t *testing.T) {
	n := EventName(MyTestEvent{}, ".")
	if n != "event_dispatcher.my.test.event" {
		t.Errorf("Wrong name. Expected '%s', got '%s'", "event_dispatcher.my.test.event", n)
	}

	n = EventName(&MyTestEvent{}, ".")
	if n != "event_dispatcher.my.test.event" {
		t.Errorf("Wrong name. Expected '%s', got '%s'", "event_dispatcher.my.test.event", n)
	}
}

type MyTestEvent struct {
}
