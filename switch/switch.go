package main

import "fmt"

func main() {
	vars := []interface{}{"Avocode", func(a int) int { return a * 2 }, true, 34}
	for _, v := range vars {
		switch t := v.(type) {
		case string:
			fmt.Printf("String with value %q\n", t)
		case int:
			fmt.Printf("Number with value %d\n", t)
		case bool:
			fmt.Printf("Boolean with value %t\n", t)
		default:
			fmt.Printf("Unknown type %T\n", t)
		}
	}
}
