package main

import (
	"fmt"
	"time"
)

func write(c chan int, x int) {
	c <- x
	close(c)
}

func main() {
	c := make(chan int)

	//Writing to channel
	go write(c, 5)

	// Reading from channel
	time.Sleep(1 * time.Second)
	fmt.Printf("%d\n", <-c)
}
