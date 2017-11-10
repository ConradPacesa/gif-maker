package main

import (
	"net/http"

	"./handlers"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}
