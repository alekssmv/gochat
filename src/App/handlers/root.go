package handlers

import (
	"net/http"
)

// Корневая страница, редирект на логин страницу
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}
