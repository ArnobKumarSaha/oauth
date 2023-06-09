package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Gitea struct {
	Client ClientDetails
}

var _ HandlerGetter = &Gitea{}

func (g *Gitea) requestForAccessToken(code string) (*http.Response, error) {
	httpClient := http.Client{}

	data := map[string]string{
		"client_id":     g.Client.ID,
		"client_secret": g.Client.Secret,
		"code":          code,
		"grant_type":    "authorization_code",
		"redirect_uri":  "http://localhost:8080/oauth/redirect", // same uri in the index.html file
	}
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reqURL := fmt.Sprintf("http://localhost:3000/login/oauth/access_token")
	fmt.Printf("reqURL = %s \n", reqURL)
	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("accept", "application/x-www-form-urlencoded")

	return httpClient.Do(req)
}

func (g *Gitea) decodeResponseBodyToGetToken(w http.ResponseWriter, res *http.Response) {
	type GiteaResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}
	var t GiteaResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// curl https://api.Gitea.com/user/repos -H "Authorization: Bearer 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"
	fmt.Printf("Gitea Response : %+v \n", t)

	w.Header().Set("Location", "/hello?access_token="+t.AccessToken)
	w.WriteHeader(http.StatusFound)
}

func (g *Gitea) usingAccessToken(token string) (*http.Response, error) {
	httpClient := http.Client{}
	/*
		curl "http://localhost:3000/test/first"     -H "accept: application/json"     -H "Authorization: bearer <TOKEN>"     -H "Content-Type: application/json"

		For more details: https://try.gitea.io/api/swagger#
	*/
	req, err := http.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/user", nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	return httpClient.Do(req)
}
