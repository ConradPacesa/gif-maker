package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"../config"
)

// Index handles the '/' request
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// process form submission
		r.ParseMultipartForm(32 << 20)
		files := r.MultipartForm.File["myfiles"]
		for _, fheader := range files {
			file, err := fheader.Open()
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			// Get file name
			fname := fheader.Filename
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			path := filepath.Join(dir, "gifs", "pics", fname)
			newFile, err := os.Create(path)
			if err != nil {
				fmt.Println(err)
			}
			defer newFile.Close()
			file.Seek(0, 0)
			io.Copy(newFile, file)
		}

	}
	config.TPL.ExecuteTemplate(w, "index.html", nil)
}
