package main

import (
	"io"
	"os"
)

func main() {
	arguments := os.Args
	var mystring string
	if len(arguments) == 1 {
		mystring = "Please provide an argument"
	} else {
		mystring = arguments[1]
	}

	io.WriteString(os.Stdout, mystring+"\n")
}
