package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	srv, err := server.NewServer(logger)
	if err != nil {
		log.Fatalf("Error create server: %v", err)
	}

	if err := srv.StartServer(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error start server: %v", err)
	}

	select {}
}
