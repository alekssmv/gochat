package handlers

import (
	"net/http"
	"gochat/session"
)

// Logout
func HandleLogout(w http.ResponseWriter, r *http.Request) {

	// Get session
	session, err := session.Store.Get(r, session.Name)
	if err != nil {
		http.Error(w, err.Error()+"Error getting session", http.StatusInternalServerError)
		return
	}

	// Delete session
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error()+"Error saving session", http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged out"))
}