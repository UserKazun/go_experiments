package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "World")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	err := server.ListenAndServe()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		return
	}
}
