package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetHelloHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello URL: ", r.URL.String())
		token := readFromURL(r, "access_token")
		res, err := requestUsingAccessToken(token)
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
func requestUsingAccessToken(token string) (*http.Response, error) {
	httpClient := http.Client{}
	/*
		curl -H "Accept: application/vnd.github+json" -H "Authorization: Bearer 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" \
		     -H "X-GitHub-Api-Version: 2022-11-28" https://api.github.com/user

		For more details: https://docs.github.com/en/rest/users
	*/
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		panic(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	return httpClient.Do(req)
}
