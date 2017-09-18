package main

import "fmt"

// Define interfaces
type Event struct {
	Data string
}

type Observer interface {
	On(Event)
}

type Notifier interface {
	Register(Observer)
	DeRegister(Observer)
	Send(Event)
}

// Define observer implementation
type stringObserver struct {
	id string
}

func (o stringObserver) On(e Event) {
	fmt.Printf("Observed [%s] event msg [%s] changed\n", o.id, e.Data)
}

// Define emitter implementation
func NewMsgEmitter() MsgEmitter {
	return MsgEmitter{
		emitters: map[Observer]struct{}{},
	}
}

type MsgEmitter struct {
	emitters map[Observer]struct{}
}

func (m MsgEmitter) Register(obj Observer) {
	m.emitters[obj] = struct{}{}
}

func (m MsgEmitter) DeRegister(obj Observer) {
	delete(m.emitters, obj)
}

func (m MsgEmitter) Send(e Event) {
	// brocase to existing emitters
	for emitter := range m.emitters {
		emitter.On(e)
	}
}

func main() {
	var notifier Notifier = NewMsgEmitter()

	receiver1 := stringObserver{"Receiver1"}
	notifier.Register(receiver1)
	receiver2 := stringObserver{"Receiver2"}
	notifier.Register(receiver2)

	notifier.Send(Event{"Hi Hi"})
	fmt.Println("====de-register receiver 2====")
	notifier.DeRegister(receiver2)

	notifier.Send(Event{"Hello"})
}
