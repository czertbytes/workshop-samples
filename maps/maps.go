package main

import "fmt"

type Person struct {
	Name, Position string
}

type ID uint

var m = map[ID]Person{
	1: Person{"Daniel Hodan", "Gopher"},
	2: Person{"Joe Doe", "React Developer"},
}

func main() {
	for k, v := range m {
		fmt.Printf("%d: %s\n", k, v.Name)
	}
}
