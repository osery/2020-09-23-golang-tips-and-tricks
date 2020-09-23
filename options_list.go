package main

import (
	"time"
)

// START1 OMIT
type config struct {
	strict    bool          // HL
	message   string        // HL
	delay     time.Duration // HL
	onSuccess func()        // HL
	onError   func(error)   // HL
}

type Option func(*config)

// START3 OMIT
func Foo(options ...Option) error { // HL
	// END1 OMIT
	c := config{
		strict:  true,              // HL
		delay:   time.Hour,         // HL
		message: "default message", // HL
	}
	for _, o := range options {
		o(&c) // HL
	}
	// Do the work according to "c" ...
	return nil
}

// END3 OMIT

// START2 OMIT
func Message(message string) Option {
	return func(c *config) { c.message = message }
}

func Strict() Option {
	return func(c *config) { c.strict = true }
}

func Delay(delay time.Duration) Option {
	return func(c *config) { c.delay = delay }
}

func OnError(f func(error)) Option {
	return func(c *config) { c.onError = f }
}

func OnSuccess(f func()) Option {
	return func(c *config) { c.onSuccess = f }
}

// END2 OMIT

func main() {
	// START4 OMIT
	err := Foo(
		Message("same message"), // HL
		Delay(time.Minute),      // HL
	)
	// END4 OMIT
	if err != nil {
		panic(err)
	}
}
