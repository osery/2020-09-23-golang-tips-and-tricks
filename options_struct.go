package main

import (
	"time"
)

// START1 OMIT
type Options struct {
	Strict    bool          // HL
	Message   string        // HL
	Delay     time.Duration // HL
	OnSuccess func()        // HL
	OnError   func(error)   // HL
}

func FooStruct(options Options) error {
	// Do the work according to "options" ...
	return nil
}

// END1 OMIT

func main() {
	// START2 OMIT
	err := FooStruct(Options{
		Message: "some text",
		Delay:   time.Minute,
	})
	// END2 OMIT
	if err != nil {
		panic(err)
	}
}
