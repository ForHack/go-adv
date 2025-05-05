package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("Starting server on :8081")
	http.ListenAndServe(":8081", nil)
}
