package main

import (
	"gochat/handlers"
	"net/http"
	_ "path/filepath"
	_ "strings"
)

// Регистрация маршрутов
func RegisterRoutes(mux *http.ServeMux, fs http.Handler) {

	// Hello world маршрут, для теста
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/", handlers.HandleRoot(fs))
	// Frontend маршруты
	mux.Handle("/register/", http.StripPrefix("/register", fs))
	mux.Handle("/login/", http.StripPrefix("/login", fs))
	mux.Handle("/contacts/", http.StripPrefix("/contacts", fs))

	// Api маршруты
	apiPrefix := "/api"
	mux.HandleFunc(apiPrefix+"/submit-register", handlers.HandleSubmitRegister)
}
