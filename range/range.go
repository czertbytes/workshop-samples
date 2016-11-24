package main

import "fmt"

func main() {
	for pos, char := range "Avocode" {
		fmt.Printf("[%d] = %c\n", pos, char)
	}
}
