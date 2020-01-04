package main

import "fmt"

type square struct {
	x int
}

type circle struct {
	x int
}

type rectangle struct {
	x int
	y int
}

func tellType(x interface{}) {
	switch x.(type) {
	case square:
		fmt.Println("It's a square")
	case circle:
		fmt.Println("It's a circle")
	case rectangle:
		fmt.Println("It's a rectangle")
	default:
		fmt.Println("Unknown type")
	}

}

func main() {
	x := square{3}
	y := circle{5}
	z := rectangle{7, 9}

	tellType(x)
	tellType(y)
	tellType(z)
}
