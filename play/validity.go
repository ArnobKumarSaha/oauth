package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	JWT_KEY      = []byte("nice_key")
	TOKEN_STRING string
)

func gen() {
	expirationTime := time.Now().Add(30 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  "ami",
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "arnob",
	})

	// Create the JWT string
	var err error
	TOKEN_STRING, err = token.SignedString(JWT_KEY)
	if err != nil {
		fmt.Println("err => ", err)
		return
	}
	fmt.Printf("Token string = %s \n", TOKEN_STRING)
}

func main() {
	gen()
	token, err := jwt.Parse(TOKEN_STRING, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the key for verification
		return JWT_KEY, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}

	// Check if the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token is valid")
	} else {
		fmt.Println("Token is invalid")
	}
}
