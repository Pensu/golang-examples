package main

import "fmt"

type twoInt struct {
	x int
	y int
}

func (a twoInt) sum(b twoInt) twoInt {
	temp := twoInt{x: a.x + b.x, y: a.y + b.y}
	return temp
}

func main() {
	a := twoInt{2, 3}
	b := twoInt{3, 4}
	fmt.Println(a.sum(b))
	fmt.Println(b.sum(a))
}
