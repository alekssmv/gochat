package handlers

import (
	"gochat/session"
	"net/http"
	"gochat/services"
)

// Logout
func HandleLogout(w http.ResponseWriter, r *http.Request) {

	jsonSuccess := services.JsonSuccessResponse

	session, _ := session.Store.Get(r, session.SessionName)

	// Delete session
	session.Options.MaxAge = -1

	// Save session
	session.Save(r, w)

	// Write json response
	jsonSuccess(w, "Logout successful", http.StatusOK)
}
