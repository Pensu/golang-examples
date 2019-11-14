package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a")
	defer func() {
		c := recover()
		if c != nil {
			fmt.Println("Handling the panic")
		}
	}()
	b()
	fmt.Println("This wont print")
}

func b() {
	panic("This is b panicking")
}

func main() {
	a()
	fmt.Println("This will print")
}
