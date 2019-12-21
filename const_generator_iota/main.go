package main

import "fmt"

func main() {

	const (
		p4_0 int = 1 << iota
		_
		p4_1
		_
		p4_2
		_
		p4_3
	)

	fmt.Println(p4_0, p4_1, p4_2, p4_3)

}
