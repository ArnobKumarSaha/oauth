package handlers

import (
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

	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		panic(err)
	}
	// We set this header since we want the response
	// as JSON
	req.Header.Set("accept", "application/json")

	return httpClient.Do(req)
}

func decodeResponseBodyToGetToken(w http.ResponseWriter, res *http.Response) {
	type OAuthAccessResponse struct {
		AccessToken string `json:"access_token"`
	}

	var t OAuthAccessResponse
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
