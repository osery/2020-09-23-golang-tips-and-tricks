package main

import (
	"fmt"
)

// START OMIT
type MyStruct struct {
	x int
	y string
}

var (
	i   int      // HL
	str string   // HL
	p   *string  // HL
	ms  MyStruct // HL
)

// END OMIT

func main() {

	fmt.Printf("i == %#v\n", i)
	fmt.Printf("str == %#v\n", str)
	fmt.Printf("p == %#v\n", p)
	fmt.Printf("ms == %#v\n", ms)
}
