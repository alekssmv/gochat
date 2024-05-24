package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	RegisterRoutes(mux)
	http.ListenAndServe(":8080", mux)
}
