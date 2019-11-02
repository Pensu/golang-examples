package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	// The first argument i.e. args[0], is the program name itself
	if len(args) == 1 {
		fmt.Println("Please provide some arguments")
		// This exits the program
		os.Exit(1)
	}

	min, _ := strconv.ParseFloat(args[1], 64)
	max, _ := strconv.ParseFloat(args[1], 64)

	for i := 2; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
