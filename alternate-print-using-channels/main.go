package main 

import (
	"fmt"
	"time"
)

func printA(in chan int) {
	for z := range in {
		fmt.Println("a: ", z)
		if z == 10 {
			close(in)
			return
		}
		in <- z+1
		time.Sleep(1*time.Second)
	}
}

func printB(in chan int) {
	for x := range in {
		fmt.Println("b: ", x)
		in <- x + 1
		time.Sleep(1*time.Second)
	}
}

func initialize (in chan int) {
	in <- 0
}

func main() {
	A := make(chan int)
	
	go initialize(A)
	go printA(A)
	time.Sleep(1*time.Second)
	printB(A)
}