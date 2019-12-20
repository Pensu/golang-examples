package main

import (
	"fmt"
	"sort"
)

// Person defines a person
type Person struct {
	name   string
	age    int
	weight int
}

func main() {
	slc := make([]Person, 0)
	slc = append(slc, Person{"Amit", 23, 48})
	slc = append(slc, Person{"Mahesh", 26, 22})
	slc = append(slc, Person{"Rakesh", 40, 89})
	slc = append(slc, Person{"Prakash", 32, 67})

	fmt.Println(slc)

	sort.Slice(slc, func(i, j int) bool {
		return slc[i].age < slc[j].age
	})
	fmt.Println(slc)

	sort.Slice(slc, func(i, j int) bool {
		return slc[i].weight < slc[j].weight
	})
	fmt.Println(slc)
}
