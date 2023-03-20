package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Github struct {
}

var _ HandlerGetter = &Github{}

func (gh *Github) requestForAccessToken(code string) (*http.Response, error) {
	httpClient := http.Client{}

	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)

	fmt.Printf("reqURL = %s \n", reqURL)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	return httpClient.Do(req)
}

func (gh *Github) decodeResponseBodyToGetToken(w http.ResponseWriter, res *http.Response) {
	type GithubResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	var t GithubResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// curl https://api.github.com/user/repos -H "Authorization: Bearer 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"
	fmt.Println("AccessToken : ", t.AccessToken)

	w.Header().Set("Location", "/hello?access_token="+t.AccessToken)
	w.WriteHeader(http.StatusFound)
}

func (gh *Github) usingAccessToken(token string) (*http.Response, error) {
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
