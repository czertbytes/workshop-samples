package main

import "fmt"

func fnPrint(fn func() (string, string)) {
	a, b := fn()
	fmt.Printf("%s\n%s\n", a, b)
}

func main() {
	a := "First Gophers"
	b := "Second Avocode"
	fnPrint(func() (string, string) {
		return b, a
	})
	fnPrint(func() (string, string) {
		return a, ""
	})
}
