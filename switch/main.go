package main

import "fmt"

func evaluate(n int) {
	switch {
	case n > 0:
		fmt.Println("Positive Integer")
	case n < 0:
		fmt.Println("Negative Integer")
	default:
		fmt.Println("Zero")
	}
}

func main() {
	evaluate(5)
	evaluate(-4)
	evaluate(0)
}
