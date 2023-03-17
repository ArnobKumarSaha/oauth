package main

import (
	"fmt"
	"github.com/Arnobkumarsaha/oauth/handlers"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Create a new redirect route route
	http.HandleFunc("/oauth/redirect", handlers.GetRedirectHandler())
	http.HandleFunc("/hello", handlers.GetHelloHandler())

	fmt.Println("Listening to :8080")
	_ = http.ListenAndServe(":8080", nil)
}
