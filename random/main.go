package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	min := 0
	max := 100
	total := 100
	seed := time.Now().Unix()

	rand.Seed(seed)
	for i := 0; i < total; i++ {
		myrand := random(min, max)
		fmt.Print(myrand)
		fmt.Print(" ")
	}
	fmt.Println()
}
