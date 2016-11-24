package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var (
		b        = true
		i uint64 = 1<<64 - 1
		z        = cmplx.Sqrt(-5 + 12i)
	)
	s := "Avocode and Gophers"

	fmt.Printf("b has type %T and value %t\n", b, b)
	fmt.Printf("i has type %T and value %d\n", i, i)
	fmt.Printf("z has type %T and value %f\n", z, z)
	fmt.Printf("s has type %T and value %s\n", s, s)
}
