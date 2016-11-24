package main

import "fmt"

type Person struct {
	Name, Position string
}

func (p Person) String() string {
	return fmt.Sprintf("Person: %s, %s", p.Name, p.Position)
}

func main() {
	persons := []Person{
		Person{"Daniel Hodan", "Gopher"},
		Person{"Joe Doe", "React Developer"},
	}

	for _, p := range persons {
		fmt.Println(p)
	}
}
