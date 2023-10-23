package event_dispatcher

// Listener defines a listener
type Listener interface {
	// Handle should accept an event and return an event and possible error
	Handle(event Event) (Event, error)
}

// ListenerFunc is a wrapper for listener as functions and implements the Listener interface
type ListenerFunc func(event Event) (Event, error)

// Handle is the implementation for the Listener interface
func (f ListenerFunc) Handle(event Event) (Event, error) {
	return f(event)
}
