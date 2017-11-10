package handlers

import (
	"net/http"

	"../config"
)

// Index handles the '/' request
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	config.TPL.ExecuteTemplate(w, "index.html", nil)
}
