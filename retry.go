package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/avast/retry-go"
)

// START1 OMIT
func tryRoll6() error {
	roll := 1 + rand.Int()%6
	fmt.Printf("Rolled: %d\n", roll)
	if roll != 6 {
		return fmt.Errorf("rolled %d, not 6", roll)
	}
	return nil
}

// END1 OMIT

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// START2 OMIT
	err := retry.Do(  // HL
		tryRoll6,
		retry.Attempts(5),          // HL
		retry.Delay(time.Second),           // HL
		retry.DelayType(retry.FixedDelay),  // HL
	)
	if err != nil {
		panic("Failed to roll a 6 in 5 attempts.")
	}
	fmt.Println("Success.")
	// END2 OMIT
}
