package main

import "fmt"

// The bridge pattern is a design pattern used in software engineering
// that is meant to "decouple an abstraction from its implementation so that the two can vary independently",
// introduced by the Gang of Four.[1]
// The bridge uses encapsulation, aggregation, and can use inheritance to separate responsibilities into different classes.

// Bridge pattern:

// 1. It is a structural pattern
// 2. Abstraction and implementation are not bound at compile time
// 3. Abstraction and implementation - both can vary without impact in client
// 4. Uses composition over inheritance.

type driverCommand struct {
	State string
}

type deviceOp interface {
	Execute(driverCommand) error
}

type driverImpl interface {
	open() error
	executeCmd(driverCommand) error
	close() error
}

type videoDriver struct{}

func (v videoDriver) open() error {
	fmt.Println("Open video driver")
	return nil
}
func (v videoDriver) executeCmd(cmd driverCommand) error {
	fmt.Printf("video driver execute: %s\n", cmd.State)
	return nil
}
func (v videoDriver) close() error {
	fmt.Println("Close video driver")
	return nil
}

type audioDriver struct{}

func (a audioDriver) open() error {
	fmt.Println("Open audio driver")
	return nil
}

func (a audioDriver) executeCmd(cmd driverCommand) error {
	fmt.Printf("audio driver execute: %s\n", cmd.State)
	return nil
}

func (a audioDriver) close() error {
	fmt.Println("Close audio driver")
	return nil
}

// work as the bridge to decouple the driver's implementation
type Device struct {
	driver driverImpl
}

func (d Device) Execute(cmd driverCommand) error {
	if err := d.driver.open(); err != nil {
		return err
	}
	if err := d.driver.executeCmd(cmd); err != nil {
		return err
	}
	return d.driver.close()
}

func DeviceExecutor(op deviceOp, cmd driverCommand) error {
	return op.Execute(cmd)
}

func main() {
	screen := Device{videoDriver{}}
	speaker := Device{audioDriver{}}

	DeviceExecutor(screen, driverCommand{
		"Turn up the brightness",
	})

	DeviceExecutor(speaker, driverCommand{
		"Turn off",
	})
}
