package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/gorilla/mux"
)

// Server — структура нашего сервера
type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// NewServer создает и возвращает новый экземпляр сервера
func NewServer(logger *log.Logger) (*Server, error) {
	router := mux.NewRouter()

	// Регистрируем наши хэндлеры
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: server,
	}, nil
}
