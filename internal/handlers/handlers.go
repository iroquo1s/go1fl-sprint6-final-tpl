package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "./index.html"

	f, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		log.Printf("Fiel not Found: %v", err)

		return
	}
	defer f.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = io.Copy(w, f)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Internal Server Error: %v", err)
		return
	}

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Error size file", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil || file == nil {
		http.Error(w, "File not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error read file", http.StatusInternalServerError)
		return
	}

	convertedStr, err := service.DetectContent(data)
	if err != nil {
		http.Error(w, "Error conver string", http.StatusInternalServerError)
		return
	}

	filename := fmt.Sprintf("%s.%s", time.Now().UTC().Format(time.RFC3339Nano), filepath.Ext(handler.Filename))

	outputFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Error create file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	if _, err := outputFile.Write([]byte(convertedStr)); err != nil {
		http.Error(w, "Error write in file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, convertedStr)
}
