package handlers

import (
	"gochat/session"
	"net/http"
)

// Logout
func HandleLogout(w http.ResponseWriter, r *http.Request) {

	session, _ := session.Store.Get(r, session.SessionName)

	// Delete session
	session.Options.MaxAge = -1

	// Save session
	session.Save(r, w)

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged out"))
}
