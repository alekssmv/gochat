package handlers

import (
	"encoding/json"
	_ "github.com/davecgh/go-spew/spew"
	"gochat/models"
	"gochat/repositories"
	"net/http"
	"gochat/session"
	"gochat/services"
)

// Создает пользователя в базе данных postgres
func HandleSubmitLogin(w http.ResponseWriter, r *http.Request) {

	jsonError := services.JsonErrorResponse
	jsonSuccess := services.JsonSuccessResponse

	if r.Method != http.MethodPost {
		jsonError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Close the request body
	defer r.Body.Close()

	// Decode JSON body into models.User struct
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if user exists in the database
	newUser, err := repositories.GetUserByUsername(user.Username)
	if err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if passwords match
	if newUser.Password != user.Password {
		jsonError(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	// Create session
	session, _ := session.Store.Get(r, session.SessionName)

	// Set session values
	session.Values["authenticated"] = true
	session.Values["username"] = user.Username

	// Save session
	session.Save(r, w)

	// response
	jsonSuccess(w, "Login successful", http.StatusOK)
}