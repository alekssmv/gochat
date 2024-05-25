package handlers

import (
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "front/index.html")
}
