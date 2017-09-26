package main

import "fmt"

// The memento pattern is a software design pattern that provides the ability to restore an object to its previous state.

type memento struct {
	value int
}

func (m *memento) GetValue() int {
	return m.value
}

type Worker struct {
	value int
}

func (w *Worker) SetValue(v int) {
	w.value = v
}

func (w *Worker) GetValue() int {
	return w.value
}

func (w *Worker) CreateMemento() *memento {
	return &memento{w.value}
}

func (w *Worker) RestoreMemento(m *memento) {
	w.value = m.GetValue()
}

// Memento
func main() {
	worker := &Worker{999}

	m := worker.CreateMemento()

	fmt.Printf("Worker value=%d\n", worker.GetValue())

	worker.SetValue(1000)

	fmt.Printf("Worker new value=%d\n", worker.GetValue())

	worker.RestoreMemento(m)

	fmt.Printf("Worker restored value=%d\n", worker.GetValue())
}
