//https://godoc.org/github.com/dgrijalva/jwt-go#example-NewWithClaims--StandardClaims
package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func main() {
	mySigningKey := []byte("AllYourBase")
	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	fmt.Printf("%v\n", ss)
}
