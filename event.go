package event_dispatcher

import (
	"fmt"
	"github.com/fatih/camelcase"
	"strings"
)

type Event interface{}

type StoppableEvent interface {
	Event
	IsPropagationStopped() bool
	WithPropagationStopped() StoppableEvent
}

func EventName(event Event, separator string) string {
	tmp := fmt.Sprintf("%T", event)
	parts := strings.Split(tmp, separator)
	tmp = parts[len(parts)-1]
	tmp = strings.ToLower(strings.Join(camelcase.Split(tmp), separator))

	return strings.TrimLeft(fmt.Sprintf("%s.%s", parts[0], tmp), "*")
}
