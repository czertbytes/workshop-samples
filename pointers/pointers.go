package main

import "fmt"

func main() {
	a := 7
	ptrA := &a
	fmt.Printf("ptrA has type %T, address %x and\nvalue %d\n", ptrA, ptrA, *ptrA)

	// b := ptrA + 3
	// will not work, ptrA is *int and 3 is int

	b := *ptrA + 3
	fmt.Printf("b: %d\n", b)
}
