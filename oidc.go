package main

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// Google OAuth2 endpoint
	googleEndpoint = google.Endpoint

	// Google OpenID Connect endpoint
	oidcEndpoint = "https://accounts.google.com"

	// Google client ID
	clientID = ""

	// Google client secret
	clientSecret = ""

	// Google scopes
	scopes = []string{oidc.ScopeOpenID, "profile", "email"}
)

func main() {
	// Create an OAuth2 config object
	oauth2Config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       scopes,
		Endpoint:     googleEndpoint,
	}

	// Create an OpenID Connect verifier
	provider, err := oidc.NewProvider(context.Background(), oidcEndpoint)
	if err != nil {
		panic(err)
	}
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	// Set up the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Redirect to Google's OAuth2 consent page
		url := oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// Get the authorization code from the request
		code := r.URL.Query().Get("code")

		fmt.Printf("code = %v \n", code)

		// Exchange the authorization code for a token
		token, err := oauth2Config.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("token = %v, %v, %v, %v  \n", token.AccessToken, token.RefreshToken, token.TokenType, token.Expiry)

		// Verify the ID token
		idToken, err := verifier.Verify(context.Background(), token.Extra("id_token").(string))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("idToken = %v \n", idToken)

		// Print the user's email address
		var val interface{}
		err = idToken.Claims(&val)
		if err != nil {
			return
		}
		fmt.Fprintf(w, "Email: %v", val)
	})

	http.ListenAndServe(":8080", nil)
}

// Copied from browser

/*
URL :
http://localhost:8080/callback?state=state
&code=4/0AVHEtk5ijvtRLq8JzfJzFWFFZ-FTTaKJ-pPpq5MiU6QCTIjOTrzYZR308hY9rZ5Fkfjzbg
&scope=email%20profile%20openid%20https://www.googleapis.com/auth/userinfo.profile%20https://www.googleapis.com/auth/userinfo.email
&authuser=0
&hd=appscode.com
&prompt=consent
*/

/*
Response :
Email: map[at_hash:wY6-T5wybwvfv9O7ep6NuQ aud:176319405093-id5048j10e7e9vc8kqo7tdt2h2kbj633.apps.googleusercontent.com azp:176319405093-id5048j10e7e9vc8kqo7tdt2h2kbj633.apps.googleusercontent.com
email:arnob@appscode.com email_verified:true exp:1.679470436e+09 family_name:Saha given_name:Arnob Kumar hd:appscode.com iat:1.679466836e+09 iss:https://accounts.google.com locale:en
name:Arnob Kumar Saha picture:https://lh3.googleusercontent.com/a/AGNmyxZIb21CYLyXnQN3GjpB8A-kkJFedYban9bcW65x=s96-c sub:105666291804906186251]
*/
