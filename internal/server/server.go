package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// Создаем новый сервер
func NewServer(logger *log.Logger) (*Server, error) {
	router := mux.NewRouter()

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

// Запускаем HTTP-сервер с выводом лога порта запуска
func (srv *Server) StartServer() error {
	srv.Logger.Printf("HTTP server started on port %s\n", srv.Server.Addr)
	return srv.Server.ListenAndServe()
}

// Остановка сервера
func (srv *Server) ShutdownServer(ctx context.Context) error {
	return srv.Server.Shutdown(ctx)
}
