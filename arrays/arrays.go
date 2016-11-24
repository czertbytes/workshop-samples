package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Gopher"
	a[1] = "Avocode"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
