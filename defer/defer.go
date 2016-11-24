package main

import "fmt"

func main() {
	fmt.Println("Fn start")
	for i := 0; i < 10; i++ {
		fmt.Printf("For %d start\n", i)
		defer fmt.Printf("Defer %d\n", i)
		fmt.Printf("For %d end\n", i)
	}
	fmt.Println("Fn end")
}
