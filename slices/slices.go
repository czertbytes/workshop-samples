package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	printSlice("a", a)
	b := make([]int, len(a))
	copy(b, a)
	b = append(b, 6)
	printSlice("b", b)
	c := b[2:4]
	printSlice("c", c)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s: len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
