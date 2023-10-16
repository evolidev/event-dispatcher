package event_dispatcher

import "strings"

type Provider interface {
	GetListenersForEvent(event Event) []Listener
}

func NewTreeProvider() *TreeProvider {
	return NewTreeProviderWithSeparator(".")
}

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

func (t *TreeProvider) Add(listenTo string, listener Listener) {
	t.listeners[listenTo] = append(t.listeners[listenTo], listener)
}

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
