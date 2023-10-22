# event-dispatcher

A simple, lightweight and flexible event dispatcher

To get started you can use the default dispatcher provided by this package or use your own an implement the `Dispatcher` interface.
The EventDispatcher takes as argument a `Provider` interface 

In this example we use our TreeProvider but you are free to give us your own provider. 
Simply implement the `Provider` interface

Below is a minimal example. 

```go
type MyTestEvent struct {
}

p := NewTreeProvider()
tlp.Add("my_package.my.test.event", ListenerFunc(func(event Event) (Event, error) {
    fmt.Println("I got called")
	
    return event, nil
}))

d := NewEventDispatcher(p)

d.Dispatch(MyTestEvent{})
```

If your listener returns an error then no other listeners will be called and the error will be reported back to the dispatcher. 

The next `Listener` will get the event returned from the previous listener.
If the Event implements `StoppableEvent` we will check  the `IsPropagationStopped` function provided by the event. 
No further listeners will be called if that function returns true. 