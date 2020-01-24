package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a web server")
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8001", nil)
}
