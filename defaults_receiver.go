package main

import (
	"fmt"
)

// START OMIT
type Struct struct {
	x int
	y string
}

func (s *Struct) operation() {
	if s == nil {
		fmt.Println("operation on nil value")
	}
}

func main() {
	var s *Struct  // HL
	s.operation()  // HL
}
// END OMIT
