package main

import "fmt"

// A flyweight is an object that minimizes memory usage by sharing as much data as possible with other similar objects;
// it is a way to use objects in large numbers when a simple repeated representation would use an unacceptable amount of memory.
// Often some parts of the object state can be shared, and it is common practice to hold them in external data structures and pass them to the objects temporarily when they are used.

const (
	OpenState = iota
	ClosedState
)

type device struct {
	info *deviceInfo
}

func (d device) String() string {
	return fmt.Sprintf("device:\n%v", d.info)
}

// save those shared state and version in deviceInfo as private field
type deviceInfo struct {
	state           *string
	firmWareVersion *string
}

func (i deviceInfo) String() string {
	return fmt.Sprintf("state:[%s], firmwareVersion:[%s]", *i.state, *i.firmWareVersion)
}

// use pool to save shared sate and version
type deviceFactory struct {
	statePool   map[int]*string
	versionPool map[string]*string
}

func NewDeviceFactory() *deviceFactory {
	return &deviceFactory{
		statePool:   map[int]*string{},
		versionPool: map[string]*string{},
	}
}

func (f *deviceFactory) toState(s int) *string {
	result := "none"
	switch s {
	case OpenState:
		result = "open"
	case ClosedState:
		result = "closed"
	default:
		fmt.Println("Not supported state")
	}
	return &result
}

func (f *deviceFactory) NewDevice(state int, version string) device {
	if _, ok := f.statePool[state]; !ok {
		f.statePool[state] = f.toState(state)
	}
	if _, ok := f.versionPool[version]; !ok {
		f.versionPool[version] = &version
	}
	return device{
		info: &deviceInfo{
			state:           f.statePool[state],
			firmWareVersion: f.versionPool[version],
		},
	}
}

func main() {
	factory := NewDeviceFactory()
	d1 := factory.NewDevice(ClosedState, "v1.20.1")
	d2 := factory.NewDevice(OpenState, "v1.20.3")
	d3 := factory.NewDevice(ClosedState, "v1.20.1")
	d4 := factory.NewDevice(OpenState, "v1.21.0")

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
}
