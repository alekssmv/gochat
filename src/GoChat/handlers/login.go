package handlers

import (
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "front/login.html")
}
