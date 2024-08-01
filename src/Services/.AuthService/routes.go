package main

import (
	"gochat/handlers"
	"gochat/middlewares"
	"net/http"
	_ "path/filepath"
	_ "strings"
)

// Регистрация маршрутов
func RegisterRoutes(mux *http.ServeMux, fs http.Handler) {

	// Корневой маршрут
	mux.HandleFunc("/", handlers.HandleRoot(fs))

	// Frontend маршруты
	mux.Handle("/register/", middlewares.NotAuthMiddleware(http.StripPrefix("/register", fs)))
	mux.Handle("/login/", middlewares.NotAuthMiddleware(http.StripPrefix("/login", fs)))

	// Защищенные маршруты
	mux.Handle("/contacts/", middlewares.AuthMiddleware(http.StripPrefix("/contacts", fs)))

	// Backend маршруты
	mux.HandleFunc("/submit-register", handlers.HandleSubmitRegister)
	mux.HandleFunc("/submit-login", handlers.HandleSubmitLogin)
	mux.HandleFunc("/logout", handlers.HandleLogout)
	mux.HandleFunc("/username", handlers.HandleUsername)

	// Hello world маршрут, для теста
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}
