package main

import (
	"net/http"
	"gochat/handlers"
)

// Регистрация маршрутов
func RegisterRoutes(mux *http.ServeMux) {

	fs := http.FileServer(http.Dir("frontend/dist"))
	mux.Handle("/", fs)
	
	mux.HandleFunc("/submit-register", handlers.HandleSubmitRegister)
}