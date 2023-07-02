package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type Todo struct {
	Data []string `json:"data"`
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
	todos := Todo{Data: []string{"job", "programming"}}

	res, _ := json.Marshal(todos)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func tempTest(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	seed := time.Now().Unix()
	random := rand.New(rand.NewSource(seed))

	err := t.Execute(w, random.Intn(10) > 5)
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

	http.HandleFunc("/sample", tempTest)

	err := server.ListenAndServe()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		return
	}
}
