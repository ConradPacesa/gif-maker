package main

import (
	"net/http"

	"./handlers"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.ListenAndServe(":80", nil)
}
