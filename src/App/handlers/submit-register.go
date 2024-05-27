package handlers

import (
	"encoding/json"
	"fmt"
	_ "github.com/davecgh/go-spew/spew"
	"gochat/db"
	"net/http"
)

type User struct {
	Username  string `json:"username"`
	Password string `json:"password"`
}

// Создает пользователя в базе данных postgres
func HandleSubmitRegister(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON body into User struct
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Close the request body
	defer r.Body.Close()

	// Create a new user in the database
	dbCon, _ := db.Connect()

	_, err := dbCon.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s registered successfully", user.Username)
}
