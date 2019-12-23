package main

import "fmt"

func main() {
	square := func(n int) int {
		return n * n
	}

	double := func(n int) int {
		return n + n
	}

	fmt.Println(square(5))
	fmt.Println(double(5))
}
