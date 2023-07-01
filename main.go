package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: &handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		return
	}
}
