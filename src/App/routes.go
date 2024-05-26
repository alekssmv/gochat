package main

import (
	"net/http"
	"gochat/handlers"
)

// Регистрация маршрутов
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("/login", handlers.HandleLogin)
	mux.HandleFunc("/register", handlers.HandleRegister)
	mux.HandleFunc("/submit-register", handlers.HandleSubmitRegister)
}