package handlers

import (
	"encoding/json"
	_ "github.com/davecgh/go-spew/spew"
	"gochat/models"
	"gochat/repositories"
	"gochat/session"
	"net/http"
)

// Создает пользователя в базе данных postgres
func HandleSubmitRegister(w http.ResponseWriter, r *http.Request) {

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

	// Create the user in the database
	err := repositories.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create session
	session, _ := session.Store.Get(r, session.SessionName)

	// Set session values
	session.Values["authenticated"] = true
	session.Values["username"] = user.Username

	// Save session
	session.Save(r, w)

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged in"))
}
