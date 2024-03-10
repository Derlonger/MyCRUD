package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"modulepath/pkg/routes"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterUserRoutes(route)

	addr := ":8183" // По умолчанию слушаем на порту 8183

	// Проверяем, задан ли порт через переменные окружения
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	server := &http.Server{
		Addr:    addr,
		Handler: route,
	}

	log.Printf("Сервер запущен по адресу http://localhost%s\n", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
