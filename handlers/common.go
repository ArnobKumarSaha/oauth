package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Handler func(http.ResponseWriter, *http.Request)

type HandlerGetter interface {
	requestForAccessToken(code string) (*http.Response, error)
	decodeResponseBodyToGetToken(w http.ResponseWriter, res *http.Response)
	usingAccessToken(token string) (*http.Response, error)
}

type ClientDetails struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
}

func readFromURL(r *http.Request, key string) string {
	// First, we need to get the value of the `code` query param
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		panic(err)
	}
	return r.FormValue(key)
}

func readResponse(res *http.Response) {
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Printf(">>>>>>>>  code: %v, response: %s \n", res.StatusCode, string(data))
}

func GetRedirectHandler(h HandlerGetter) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("URL in redirect handler: ", r.URL.String())
		code := readFromURL(r, "code")
		res, err := h.requestForAccessToken(code)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()
		h.decodeResponseBodyToGetToken(w, res)
	}
}

func GetHelloHandler(h HandlerGetter) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("URL in hello handler: ", r.URL.String())
		token := readFromURL(r, "access_token")
		res, err := h.usingAccessToken(token)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not send HTTP request to github: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return
		}
		fmt.Printf("code: %v, response: %s \n", res.StatusCode, string(data))
	}
}
