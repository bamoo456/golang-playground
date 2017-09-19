package main

import "fmt"

// Similar to electrical fuses that prevent fires when a circuit that is connected to the electrical grid
// starts drawing a high amount of power which causes the wires to heat up and combust,
// the circuit breaker design pattern is a fail-first mechanism that shuts down the circuit,
// request/response relationship or a service in the case of software development, to prevent bigger failures.

// FROM: https://docs.microsoft.com/en-us/azure/architecture/patterns/circuit-breaker
// Handle faults that might take a variable amount of time to recover from, when connecting to a remote service or resource. This can improve the stability and resiliency of an application.

// existing implementation: https://github.com/sony/gobreaker

func main() {
	// TODO:
	fmt.Println("NOT IMPLEMENTED")
}
