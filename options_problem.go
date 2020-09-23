package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/avast/retry-go"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// START OMIT
	err := retry.Do(
		tryRoll6,
		retry.Attempts(5),          // HL
		retry.Delay(time.Second),           // HL
		retry.DelayType(retry.FixedDelay),  // HL
	)
	// END OMIT
	if err != nil {
		panic("Failed to roll a 6 in 5 attempts.")
	}
	fmt.Println("Success.")
}
