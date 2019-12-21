package main

import "fmt"

func main() {
	arr := [5]int{4, 5, 6, 7, 8}
	myMap := map[int]int{}

	for i, val := range arr {
		myMap[i+1] = val
	}

	fmt.Println(myMap)
}
