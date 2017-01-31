package main

import "net/http"

func main() {
	http.Handle("/", nil)
	http.ListenAndServe("127.0.0.1:8081", nil)
}
