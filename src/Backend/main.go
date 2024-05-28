package main

import (
	"net/http"
	_ "github.com/lib/pq"
)

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("dist"))
	// Регистрация маршрутов
	RegisterRoutes(mux, fs)

	http.ListenAndServe(":8080", mux)
}
