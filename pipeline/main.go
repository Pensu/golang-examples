package main

import (
	"fmt"
	"math/rand"
	"time"
)

var CLOSEA = false

var data = make(map[int]bool)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		_, ok := data[x]

		if ok {
			CLOSEA = true
		} else {
			data[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	sum := 0
	for x2 := range in {
		sum = sum + x2
	}

	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)

	go first(1, 100, A)
	go second(B, A)
	third(B)
}
