package main

import (
	"fmt"
)

func main() {
	arr := [7]int{2, 3, 4, 5, 6, 7, 8}
	searchElem := 6

	lowerBound := 0
	upperBound := len(arr) - 1

	for lowerBound <= upperBound {
		mid := (upperBound + lowerBound) / 2

		if arr[mid] == searchElem {
			fmt.Println(mid)
			break
		} else if arr[mid] > searchElem {
			upperBound = mid - 1
		} else if arr[mid] < searchElem {
			lowerBound = mid + 1
		}
	}
}
