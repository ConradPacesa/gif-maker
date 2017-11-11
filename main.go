package main

import (
	"net/http"

	"github.com/ConradPacesa/gif-maker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}
