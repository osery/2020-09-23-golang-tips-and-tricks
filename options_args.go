package main

import (
	"time"
)

// START1 OMIT
func FooArgs(
	strict bool,         // HL
	message string,      // HL
	delay time.Duration, // HL
	onSuccess func(),    // HL
	onError func(error), // HL
) error {
	// Do the work ...
	return nil
}

// END1 OMIT

func main() {
	// START2 OMIT
	err := FooArgs(false, "some text", time.Minute, nil, nil)
	// END2 OMIT
	if err != nil {
		panic(err)
	}
}
