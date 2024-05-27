package main

import (
	"net/http"
	"gochat/handlers"
)

// Регистрация маршрутов
func RegisterRoutes(mux *http.ServeMux) {

	// serve static files
	fs := http.FileServer(http.Dir("dist"))
	mux.Handle("/", fs)

	// print hello world
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	
	apiPrefix := "/api"
	mux.HandleFunc(apiPrefix+"/submit-register", handlers.HandleSubmitRegister)
}