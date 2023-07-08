package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func tempTest(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	seed := time.Now().Unix()
	random := rand.New(rand.NewSource(seed))

	err := t.Execute(w, random.Intn(10) > 5)
	if err != nil {
		return
	}
}

func experimentsHello(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "experiments_token",
		Value: "hello",
		Path:  "/",
	}

	http.SetCookie(w, &cookie)
}

func experimentsGood(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "experiments_token",
		Value: "good",
		Path:  "/",
	}

	http.SetCookie(w, &cookie)
}

func hello(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("experiments_token")
	if err != nil {
		fmt.Fprintln(w, "cannot get the cookie")
		return
	}

	// If there is GOOD, then HELLO is not shown.
	if token.Value == "hello" {
		fmt.Fprintln(w, "get cookie")
		if err != nil {
			return
		}

		return
	} else {
		fmt.Fprintln(w, "cannot get 'hello' cookie")
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/sample", tempTest)
	http.HandleFunc("/experiments/hello", experimentsHello)
	http.HandleFunc("/experiments/good", experimentsGood)
	http.HandleFunc("/hello", hello)

	err := server.ListenAndServe()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		return
	}
}
