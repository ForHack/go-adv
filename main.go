package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
}

func main() {
	router := http.NewServeMux()
	http.HandleFunc("/hello", hello)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Starting server on :8081")
	server.ListenAndServe()
}
