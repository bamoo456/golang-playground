package main

import "fmt"

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

// Strategy pattern is similar to Template pattern except in its granularity.
// Strategy pattern lets you change the guts of an object. Decorator pattern lets you change the skin.
func main() {
	var worker OperateWorker

	worker = OperateWorker{intOperation{1}}

	worker.Operation.Apply()

	worker = OperateWorker{strOperation{"string worker"}}

	worker.Operation.Apply()

}
