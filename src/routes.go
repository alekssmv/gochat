package main

import (
	"net/http"
	"GoChat/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HandleRoot)
}