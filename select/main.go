package main

import (
	"fmt"
	"time"
)

func serve1(ch chan int) {
	time.Sleep(5 * time.Second)
	ch <- 5
}

func serve2(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 3
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go serve1(a)
	go serve2(b)

	time.Sleep(10 * time.Second)

	select {
	case n1 := <-a:
		fmt.Println(n1)
	case n2 := <-b:
		fmt.Println(n2)
	}
}
