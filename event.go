package event_dispatcher

import (
	"fmt"
	"github.com/fatih/camelcase"
	"strings"
)

type Event interface{}

// StoppableEvent defines a specific kind of Event
type StoppableEvent interface {
	Event
	// IsPropagationStopped should tell the dispatcher that the propagation should be stopped
	IsPropagationStopped() bool
	// WithPropagationStopped returns a new event with the indicator that it should be stopped
	WithPropagationStopped() StoppableEvent
}

// EventName parses the name of the event struct and transform it from camelCase to lowercase string
// where each change is seperated by given separator
func EventName(event Event, separator string) string {
	tmp := fmt.Sprintf("%T", event)
	parts := strings.Split(tmp, separator)
	tmp = parts[len(parts)-1]
	tmp = strings.ToLower(strings.Join(camelcase.Split(tmp), separator))

	return strings.TrimLeft(fmt.Sprintf("%s.%s", parts[0], tmp), "*")
}
