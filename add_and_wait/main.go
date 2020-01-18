package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup

	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(x int) {
			defer wait.Done()
			fmt.Println(x)
		}(i)
	}

	wait.Wait()
	fmt.Println("Exiting...")
	fmt.Printf("%#v\n", wait)
}
