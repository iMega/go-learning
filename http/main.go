package main

import (
	"io"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, req.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", myHandler)

	err := http.ListenAndServe("0.0.0.0:8081", mux)
	log.Fatal(err)
}
