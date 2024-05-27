package handlers

import (
	"fmt"
	"net/http"
)

// Корневая страница
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleRoot")

	http.FileServer(http.Dir("frontend/dist")).ServeHTTP(w, r)
}
