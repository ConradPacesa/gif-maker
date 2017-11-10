package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ConradPacesa/gif-maker/config"
)

// Index handles the '/' request
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// process form submission
		r.ParseMultipartForm(2048)
		files := r.MultipartForm.File["myfiles"]
		for _, fheader := range files {
			go copyFiles(fheader)
		}

	}
	config.TPL.ExecuteTemplate(w, "index.html", nil)
}

func copyFiles(fh *multipart.FileHeader) {
	file, err := fh.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// Get file name
	fname := fh.Filename
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
