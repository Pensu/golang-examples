package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a OS file pointer, stdin, stdout are files
	var f *os.File
	// Using stdin file.
	f = os.Stdin
	defer f.Close()

	// Create a new scanner using stdin fd
	scanner := bufio.NewScanner(f)
	// Keep Scanning and print using Text
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
