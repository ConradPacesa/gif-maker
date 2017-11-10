package handlers

import (
	"net/http"

	"../config"
)

// Index handles the '/' request
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		config.TPL.ExecuteTemplate(w, "index.html", nil)
	}

	if r.Method == "POST" {
		// TODO take in photo
	}

}
