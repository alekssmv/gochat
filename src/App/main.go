package main

import (
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	// Register routes
	RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
