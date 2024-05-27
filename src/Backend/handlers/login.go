package handlers

import (
	"net/http"
)

// Показывает странуцу регистрации
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "front/login.html")
}
