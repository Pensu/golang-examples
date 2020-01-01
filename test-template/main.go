package main

import (
	"fmt"
	"os"
	"text/template"
)

// Data is a data struct
type Data struct {
	Number int
	Square int
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please run with template file")
		return
	}

	tFile := arguments[1]
	DATA := [][]int{{1, 1}, {2, 4}, {3, 9}, {4, 16}}

	var arr []Data

	for _, i := range DATA {
		temp := Data{Number: i[0], Square: i[1]}
		arr = append(arr, temp)
	}

	fmt.Println(arr)

	t := template.Must(template.ParseGlob(tFile))
	t.Execute(os.Stdout, arr)

}
