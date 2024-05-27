package main

import (
	"net/http"
	_ "github.com/lib/pq"
)

func main() {

	mux := http.NewServeMux()
	// Register routes
	RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
