package event_dispatcher

import "testing"

func TestName(t *testing.T) {
	n := EventName(MyTestEvent{}, ".")
	if n != "my.test.event" {
		t.Errorf("Wrong name. Expected '%s', got '%s'", "my.test.event", n)
	}
}

type MyTestEvent struct {
}
