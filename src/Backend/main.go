package main

import (
	"gochat/db"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()

	// Подключение к базе данных
	db.Connect(5, 5*time.Second)

	// Serve static files from the "dist" folder
	fs := http.FileServer(http.Dir("dist"))

	// Регистрация маршрутов
	RegisterRoutes(mux, fs)

	http.ListenAndServe(":8080", mux)
}
