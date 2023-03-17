package handlers

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	clientID     string
	clientSecret string
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")

	flag.StringVar(&clientID, "client-id", clientID, "client ID")
	flag.StringVar(&clientSecret, "client-secret", clientSecret, "client Secret")
	flag.Parse()
	fmt.Println("==>", clientID, clientSecret)

}

func GetRedirectHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request URL: ", r.URL.String())
		code := readFromURL(r, "code")
		res, err := requestForAccessToken(code)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()
		decodeResponseBodyToGetToken(w, res)
	}
}

func requestForAccessToken(code string) (*http.Response, error) {
	httpClient := http.Client{}

	// for Github
	//reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)

	data := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
		"grant_type":    "authorization_code",
		"redirect_uri":  "http://localhost:8080/oauth/redirect", // same uri in the index.html file
	}
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// For Gitea
	reqURL := fmt.Sprintf("http://localhost:3000/login/oauth/access_token")
	fmt.Printf("reqURL = %s \n", reqURL)
	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		panic(err)
	}
	// We set this header since we want the response
	// as JSON
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("accept", "application/x-www-form-urlencoded")

	return httpClient.Do(req)
}

func decodeResponseBodyToGetToken(w http.ResponseWriter, res *http.Response) {
	//
	//data, err := io.ReadAll(res.Body)
	//if err != nil {
	//	return
	//}
	//fmt.Printf(">>>>>>>>  code: %v, response: %s \n", res.StatusCode, string(data))

	//type GithubResponse struct {
	//	AccessToken string `json:"access_token"`
	//}
	//var t GithubResponse
	//if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
	//	fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

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

	// curl https://api.github.com/user/repos -H "Authorization: Bearer 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"
	fmt.Println("AccessToken : ", t.AccessToken)

	w.Header().Set("Location", "/hello?access_token="+t.AccessToken)
	w.WriteHeader(http.StatusFound)
}
