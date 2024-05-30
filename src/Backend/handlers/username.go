package handlers

import (
	_"github.com/davecgh/go-spew/spew"
	"net/http"
	"gochat/session"
)

// Получение имени пользователя
func HandleUsername(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Close the request body
	defer r.Body.Close()

	// Get session
	session, _ := session.Store.Get(r, session.SessionName)

	// Retrieve username from session
	username := session.Values["username"]

	if username == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(username.(string)))
}
