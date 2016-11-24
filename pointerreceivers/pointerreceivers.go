package main

import "fmt"

type Person struct {
	Name, Position string
}

func (p Person) String() string {
	return fmt.Sprintf("Person: %s, %s", p.Name, p.Position)
}

// START OMIT
func (p *Person) ConvertToGopherism() {
	p.Position = "Gopher"
}

func main() {
	persons := []Person{
		Person{"Daniel Hodan", "Gopher"},
		Person{"Joe Doe", "React Developer"},
	}

	for _, p := range persons {
		p.ConvertToGopherism()
		fmt.Println(p)
	}
}

// END OMIT
