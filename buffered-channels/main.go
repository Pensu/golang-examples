package main

import (
	"fmt"
)

func main() {
	a := make(chan int, 5)

	for i := 0; i < 10; i++ {
		select {
		case a <- i:
		default:
			fmt.Println("No space for number ", i)
		}
	}

	for i := 0; i < 10; i++ {
		select {
		case x := <-a:
			fmt.Println(x)
		default:
			fmt.Println("Nothing to be done")
		}
	}
}
