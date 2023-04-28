package main

import (
	"flag"
	"fmt"
	"github.com/Arnobkumarsaha/oauth/handlers"
	"net/http"
	"os"
	"strings"
)

var (
	cl         handlers.ClientDetails
	authServer string
)

func init() {
	cl.ID = os.Getenv("CLIENT_ID")
	cl.Secret = os.Getenv("CLIENT_SECRET")
	authServer = os.Getenv("AUTH_SERVER")

	flag.StringVar(&cl.ID, "client-id", cl.ID, "client ID")
	flag.StringVar(&cl.Secret, "client-secret", cl.Secret, "client Secret")
	flag.StringVar(&authServer, "auth-server", authServer, "Available values : Github, Gitea")
	flag.Parse()
	fmt.Printf("==> %v, %v, %v \n", cl.ID, cl.Secret, authServer)

}

// go run *.go --client-id=<> --client-secret=<> --auth-server=github

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// determine which type to use
	authServer = strings.ToUpper(authServer)
	var g handlers.HandlerGetter
	if authServer == "GITHUB" {
		g = &handlers.Github{Client: cl}
	} else if authServer == "GITEA" {
		g = &handlers.Gitea{Client: cl}
	} else if authServer == "B3" {
		g = &handlers.B3{Client: cl}
	}
	http.HandleFunc("/oauth/redirect", handlers.GetRedirectHandler(g))
	http.HandleFunc("/hello", handlers.GetHelloHandler(g))

	fmt.Println("Listening to :8080")
	_ = http.ListenAndServe(":8080", nil)
}
