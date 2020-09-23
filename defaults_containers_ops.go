package main

import (
	"fmt"
)

func main() {
	// START OMIT
	var (
		l []string
		m map[string]int
	)

	s := len(l)                   // HL
	a := append(l, "foo")  // HL
	v, ok := m["foo"]             // HL
	// END OMIT

	fmt.Printf("l == %#v\n", l)
	fmt.Printf("m == %#v\n", m)
	fmt.Println()
	fmt.Printf("s == %#v\n", s)
	fmt.Printf("a == %#v\n", a)
	fmt.Printf("v == %#v, ok == %#v\n", v, ok)
}
