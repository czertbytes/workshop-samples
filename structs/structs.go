package main

import "fmt"

type Person struct {
	Name, Position string
	Company
}

type Company struct {
	Name string
}

func main() {
	p := Person{
		Name:     "Daniel Hodan",
		Position: "Gopher",
		Company: Company{
			Name: "Avocode",
		},
	}
	fmt.Printf("%s works at %s\n", p.Name, p.Company.Name)
}
