package main

import (
	"net/http"
	"gochat/db"
)
func main() {
	mux := http.NewServeMux()

	// Подключение к базе данных
	db.Connect()

	// Serve static files react app
	fs := http.FileServer(http.Dir("dist"))

	// Регистрация маршрутов
	RegisterRoutes(mux, fs)

	http.ListenAndServe(":8080", mux)
}