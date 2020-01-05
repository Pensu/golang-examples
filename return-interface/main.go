package main

import "fmt"

func retValue(n int) interface{} {
	switch n {
	case 1:
		return "100"
	case 2:
		return "Hello world"
	case 10:
		return "-100"
	default:
		return "Unknown input"
	}

}

func main() {
	fmt.Println(retValue(1))
	fmt.Println(retValue(2))
	fmt.Println(retValue(10))
	fmt.Println(retValue(25))

}
