package main

import "fmt"

// In short, a proxy is a wrapper or agent object that is being called by the client to access the real serving object behind the scenes.
// Use of the proxy can simply be forwarding to the real object, or can provide additional logic.
// In the proxy, extra functionality can be provided,
// for example caching when operations on the real object are resource intensive, or checking preconditions before operations on the real object are invoked.

type Printer interface {
	Name() string
	Print() error
	SetName(string)
}

type LaserPrinter struct {
	name string
}

func (p *LaserPrinter) Name() string {
	return p.name
}

func (p *LaserPrinter) Print() error {
	fmt.Println("=======Printing .......=======")
	return nil
}

func (p *LaserPrinter) SetName(name string) {
	p.name = name
}

type PrinterProxy struct {
	printerName string
	realPrinter Printer
}

func (p *PrinterProxy) Name() string {
	return p.printerName
}

// assume the Print() is a heavy operation
func (p *PrinterProxy) Print() error {
	if p.realPrinter == nil {
		// init the real printer if it does not exist
		p.realPrinter = &LaserPrinter{p.printerName}
	} else {
		// configure the printer with latest proxy configuration
		p.realPrinter.SetName(p.printerName)
	}

	// do some operations

	fmt.Println("=====Proxy printing....=====")
	return p.realPrinter.Print()
}

func (p *PrinterProxy) SetName(name string) {
	// In the proxy, we don't realy need to configure the printer name at present
	p.printerName = name
}

func main() {
	var myPrinter Printer

	//  use real printer
	myPrinter = &LaserPrinter{"HP"}

	fmt.Println("Current printer name is ", myPrinter.Name())

	myPrinter.Print()

	fmt.Println("--------------------------")

	myPrinter = &PrinterProxy{}
	myPrinter.SetName("EPSON")

	fmt.Println("Current printer name is ", myPrinter.Name())

	myPrinter.Print()
}
