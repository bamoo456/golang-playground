package main

import "log"

// NOTE:
// * Unlike Adapter pattern, the object to be decorated is obtained by injection.
// * Decorators should not alter the interface of an object.
func LogDecorator(fn func(int) int) func(int) int {

	return func(value int) int {
		log.Println("Logging the entrance of the function")
		res := fn(value)
		log.Println("Logging the exit of the function")
		return res
	}
}

func powByTwo(value int) int { return value * value }

func main() {
	v := 10

	v2 := powByTwo(v)

	log.Println("v2 is ", v2)

	newPowByTwo := LogDecorator(powByTwo)

	v2 = newPowByTwo(v)

	log.Println("v2 is ", v2)
}
