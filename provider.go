package event_dispatcher

import "strings"

type Provider interface {
	// GetListenersForEvent is responsible to return a slice of Listener for a given Event
	GetListenersForEvent(event Event) []Listener
}

// NewTreeProvider returns a Provider which allows to listen to a set of events
// Eg. you got an Event mypackage.ModelTestEvent
// Now one can register to all events happening with "" or all events happening just in "mypackage"
// or all events of the model "mypackage.model" or explicit the event "mypackage.model.test"
// The provider will return from more global to more restrictive listeners
//
// For the example above you would get the listeners in the following order
// listeners subscribed for ""
// listeners subscribed for "mypackage"
// listeners subscribed for "mypackage.model"
// listeners subscribed for "mypackage.model.test"
// each level returns their subscriber in the order they got them
func NewTreeProvider() *TreeProvider {
	return NewTreeProviderWithSeparator(".")
}

// NewTreeProviderWithSeparator same as NewTreeProvider but specifying a custom separator is possible
func NewTreeProviderWithSeparator(separator string) *TreeProvider {
	return &TreeProvider{
		listeners: make(map[string][]Listener),
		separator: separator,
	}
}

type TreeProvider struct {
	listeners map[string][]Listener
	separator string
}

// Add a listener to a "level". "" => all events
func (t *TreeProvider) Add(listenTo string, listener Listener) {
	t.listeners[listenTo] = append(t.listeners[listenTo], listener)
}

// GetListenersForEvent implementing the interface
func (t *TreeProvider) GetListenersForEvent(event Event) []Listener {
	l := make([]Listener, 0)

	parts := strings.Split(EventName(event, t.separator), t.separator)
	nextTarget := ""

	for i := 0; i <= len(parts); i++ {
		nextTarget = strings.Join(parts[:i], t.separator)
		l = append(l, t.listeners[nextTarget]...)
	}

	return l
}
