package handlers

import (
	_"github.com/davecgh/go-spew/spew"
	"net/http"
	"gochat/session"
	"encoding/json"
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

	// Write Json response
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(map[string]string{
		"username": username.(string),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
