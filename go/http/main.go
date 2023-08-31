package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
