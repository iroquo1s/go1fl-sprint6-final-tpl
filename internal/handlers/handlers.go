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
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	filePath := "./index.html"

	http.ServeFile(w, r, filePath)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil { // Limit file size to 10 MB
		http.Error(w, "File exceeds size limit", http.StatusBadRequest)
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
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	convertedStr, err := service.DetectContent(data)
	if err != nil {
		http.Error(w, "Error converting string", http.StatusInternalServerError)
		return
	}

	filename := fmt.Sprintf("%s%s", time.Now().UTC().Format(time.RFC3339Nano), filepath.Ext(handler.Filename))

	outputFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	if _, err := outputFile.WriteString(convertedStr); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintf(w, convertedStr); err != nil {
		log.Println("Output error:", err)
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
}
