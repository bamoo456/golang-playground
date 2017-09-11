package main

import (
	"flag"
	"fmt"
	"time"
)

type Celsius float64
type Fahrenheit float64

func FToC(f Fahrenheit) Celsius {
	return (Celsius(f) - 32) * 5 / 9
}

type celsiusFlag struct{ Celsius }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	// this flag will be registered into command flag set
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

// define the command flag "-sleep"
// this "-sleep" flag will be registered into flag package
var sleepPeriod = flag.Duration("sleep", 1*time.Second, "sleep period")

// define the command flag "-temp"
var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	// parse the flag values which user key in
	// e.g. user can use command: "./flag -sleep 50ms" to control the sleep period
	//
	// NOTE: if user key in the "-temp ", the celsiusFlag.Set() will be called
	flag.Parse()

	fmt.Printf("Sleeping for %v...\n", *sleepPeriod)
	time.Sleep(*sleepPeriod)
	fmt.Printf("Current tempature is %v \n", temp.String())
}
