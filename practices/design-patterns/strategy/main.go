package main

// Strategy: ( Behavioural pattern)

// Strategy patterns enable you to switch between multiple algorithms from a family of algorithms at run time.

// Use Strategy pattern when :

// Multiple versions of algorithms are required
// The behaviour of class has to be changed dynamically at run time
// Avoid conditional statements

// REFERENCE:
// https://stackoverflow.com/questions/370258/real-world-example-of-the-strategy-pattern/35180265#35180265

import (
	"errors"
	"fmt"
)

type Operation interface {
	Apply() error
}

type intOperation struct {
	data int
}

func (i intOperation) Apply() error {
	fmt.Println("Now using intOperation with data=", i.data)
	return nil
}

type strOperation struct {
	name string
}

func (str strOperation) Apply() error {
	fmt.Println("Now using strOperation with name=", str.name)
	return nil
}

type OperateWorker struct {
	Operation
}

// Apply the worker in runtime (choose the strategy in runtime)
func workerApllier(c interface{}) error {
	var worker OperateWorker
	switch c.(type) {
	case int:
		worker = OperateWorker{intOperation{c.(int)}}
	case string:
		worker = OperateWorker{strOperation{c.(string)}}
	default:
		fmt.Println("Not abble to apply worker")
		return errors.New("Not supported condition")
	}
	return worker.Apply()
}

// Strategy pattern is similar to Template pattern except in its granularity.
// Strategy pattern lets you change the guts of an object. Decorator pattern lets you change the skin.
func main() {
	workerApllier(1)

	workerApllier("Hello world")

	// Not supported condition
	workerApllier(1.23)
}
