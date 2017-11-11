package handlers

import (
	"fmt"
	"image"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ConradPacesa/gif-maker/config"
)

// Index handles the '/' request
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// process form submission
		r.ParseMultipartForm(2048)
		files := r.MultipartForm.File["myfiles"]
		for _, fheader := range files {
			copyFiles(fheader)
		}
		convertToGif()
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

func convertToGif() {
	files := []string{}

	dir, err := os.Getwd()
	searchDir := filepath.Join(dir, "gifs", "pics")
	if err != nil {
		fmt.Println(err)
	}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	//fmt.Println(files)

	gifFiles := []string{}
	for i, name := range files[2:] {
		f, err := os.Open(name)
		if err != nil {
			fmt.Printf("There was an error opening the file: %v", err)
		}
		t, _, err := image.Decode(f)
		if err != nil {
			fmt.Printf("There was an error decoding the image: %v", err)
		}

		nm := strconv.Itoa(i)
		fn := fmt.Sprintf(nm) + ".gif"
		f, _ = os.Create(filepath.Join(dir, "gifs", fn))
		gif.Encode(f, t, nil)
		gifFiles = append(gifFiles, filepath.Join(dir, "gifs", fn))
	}

	//fmt.Println(gifFiles)
	outGif := &gif.GIF{}
	for _, name := range gifFiles {
		f, _ := os.Open(name)
		inGif, _ := gif.Decode(f)
		f.Close()

		outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 0)
	}

	gifPath := filepath.Join(dir, "public", "gif", "output.gif")
	f, _ := os.OpenFile(gifPath, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, outGif)
}
