package main

import "fmt"

func eval(i int) (string, error) {
	if i%2 == 0 {
		return fmt.Sprintf("*** %.2d", i), nil
	} else if i == 5 {
		return fmt.Sprintf("%b ###", i), nil
	}

	return "", fmt.Errorf("Ugly number")
}

func main() {
	for i := 1; i < 10; i++ {
		if res, err := eval(i); err == nil {
			fmt.Printf("%d => %s\n", i, res)
		}
	}
}
