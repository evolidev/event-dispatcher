package event_dispatcher

type Dispatcher interface {
	Dispatch(event Event) error
}

func NewEventDispatcher(provider Provider) *EventDispatcher {
	return &EventDispatcher{provider: provider}
}

type EventDispatcher struct {
	provider Provider
}

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

func NewGoDispatcher(provider Provider) Dispatcher {
	return &GoDispatcher{dispatcher: NewEventDispatcher(provider)}
}

type GoDispatcher struct {
	dispatcher *EventDispatcher
}

func (d *GoDispatcher) Dispatch(event Event) error {
	go d.dispatcher.Dispatch(event)

	return nil
}
