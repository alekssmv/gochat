package handlers

import (
	"net/http"
)

// Показывает странуцу логина
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "front/register.html")
}
