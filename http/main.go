package main

import (
	"io"
	"log"
	"net/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, req.URL.Path)
}
type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", myHandler)

	err := http.ListenAndServe("0.0.0.0:8081", mux)
	log.Fatal(err)

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
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
}
