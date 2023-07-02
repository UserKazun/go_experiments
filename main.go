package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Title string `json:"title"`
}

func APIAddTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func APIGetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todos []Todo
	todo1 := Todo{Title: "job"}
	todos = append(todos, todo1)

	todo2 := Todo{Title: "programming"}
	todos = append(todos, todo2)

	res, _ := json.Marshal(todos)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/api/todo", APIAddTodo)
	http.HandleFunc("/api/todos", APIGetTodo)

	err := server.ListenAndServe()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		return
	}
}
