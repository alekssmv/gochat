package handlers

import (
	"encoding/json"
	_ "github.com/davecgh/go-spew/spew"
	"gochat/models"
	"net/http"
	"gochat/repositories"
	"gochat/session"
)

// Создает пользователя в базе данных postgres
func HandleSubmitLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Close the request body
	defer r.Body.Close()

	// Decode JSON body into models.User struct
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if user exists in the database
	newUser, err := repositories.GetUserByUsername(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if passwords match
	if newUser.Password != user.Password {
		http.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	// Create session
	session, _ := session.Store.Get(r, session.Name)
    session.Values["authenticated"] = true
    session.Values["username"] = user.Username
    session.Save(r, w)

	
	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged in"))
}