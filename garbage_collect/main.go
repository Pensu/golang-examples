package main

import (
	"fmt"
	"runtime"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println("Number of GC: ", mem.NumGC)
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000)
		if s == nil {
			fmt.Println("Error")
		}
		printStats(mem)
	}
}
