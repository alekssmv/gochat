package handlers

import (
	_ "github.com/davecgh/go-spew/spew"
	"gochat/services"
	"gochat/session"
	"net/http"
)

// Получение имени пользователя
func HandleUsername(w http.ResponseWriter, r *http.Request) {

	jsonError := services.JsonErrorResponse
	json := services.JsonResponse

	if r.Method != http.MethodGet {
		jsonError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Close the request body
	defer r.Body.Close()

	// Get session
	session, _ := session.Store.Get(r, session.SessionName)

	// Retrieve username from session
	username := session.Values["username"]

	if username == nil {
		jsonError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Write Json response
	json(w, map[string]string{"username": username.(string)}, http.StatusOK)
}
