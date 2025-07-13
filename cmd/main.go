package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создадим логгер
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	// Получим экземпляр сервера
	srv, err := server.NewServer(logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Сервер создан")

	// Запускаем сервер
	err = srv.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}

}
