package event_dispatcher

// Dispatcher interface only has the Dispatch function
type Dispatcher interface {
	Dispatch(event Event) error
}

// NewEventDispatcher returns a new EventDispatcher
func NewEventDispatcher(provider Provider) *EventDispatcher {
	return &EventDispatcher{provider: provider}
}

// EventDispatcher implements the Dispatcher interface and handles a default dispatching logic
type EventDispatcher struct {
	provider Provider
}

// Dispatch a given event to all listeners provided by the listener
// If a listener returns an error no more listener gets called
// A listener gets the event returned by the previous listener
// if an Event is a type of StoppableEvent then no more listener get called
func (d *EventDispatcher) Dispatch(event Event) error {
	var err error
	for _, listener := range d.provider.GetListenersForEvent(event) {
		event, err = listener.Handle(event)

		if err != nil {
			return err
		}

		if stoppable, ok := event.(StoppableEvent); ok {
			stoppable.IsPropagationStopped()

			return nil
		}
	}

	return nil
}

// NewGoDispatcher returns an instance of Dispatcher
// Internally it uses the EventDispatcher
func NewGoDispatcher(provider Provider) Dispatcher {
	return &GoDispatcher{dispatcher: NewEventDispatcher(provider)}
}

type GoDispatcher struct {
	dispatcher *EventDispatcher
}

// Dispatch uses the EventDispatcher Dispatch method
// but runs it in a go routine
func (d *GoDispatcher) Dispatch(event Event) error {
	go d.dispatcher.Dispatch(event)

	return nil
}
