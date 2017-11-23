package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	mySigningKey = "WOW,MuchShibe,ToDogge"
)

func main() {
	createdToken, err := ExampleNew([]byte(mySigningKey))
	if err != nil {
		fmt.Println("Creating token failed")
	}
	fmt.Println(createdToken)
	ExampleParse(createdToken, mySigningKey)
	ExampleParse(createdToken + "1", mySigningKey)
}

func ExampleNew(mySigningKey []byte) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["foo"] = "bar"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}

func ExampleParse(myToken string, myKey string) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})

	fmt.Println(token)

	if err == nil && token.Valid {
		fmt.Println("Your token is valid.  I like your style.")
	} else {
		fmt.Println("This token is terrible!  I cannot accept this.")
	}
}