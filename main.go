package main

import (
	"fmt"
	"github.com/Arnobkumarsaha/oauth/handlers"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	g := &handlers.Github{}
	http.HandleFunc("/oauth/redirect", handlers.GetRedirectHandler(g))
	http.HandleFunc("/hello", handlers.GetHelloHandler(g))

	fmt.Println("Listening to :8080")
	_ = http.ListenAndServe(":8080", nil)
}
