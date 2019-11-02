package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide some arguments")
		os.Exit(1)
	}

	sum, _ := strconv.ParseFloat(args[1], 64)
	for i := 2; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[i], 64)
		sum = sum + n
	}

	fmt.Println(sum)
}
