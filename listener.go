package event_dispatcher

type Listener interface {
	Handle(event Event) (Event, error)
}

type ListenerFunc func(event Event) (Event, error)

func (f ListenerFunc) Handle(event Event) (Event, error) {
	return f(event)
}
