package handlers

import (
	"fmt"
	"net/http"
	"os"
)

type Handler func(http.ResponseWriter, *http.Request)

func readFromURL(r *http.Request, key string) string {
	// First, we need to get the value of the `code` query param
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		panic(err)
	}
	return r.FormValue(key)
}
