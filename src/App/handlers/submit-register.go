package handlers

import (
	"net/http"
	"gochat/db"
)

// Создает пользователя в базе данных postgres
func HandleSubmitRegister(w http.ResponseWriter, r *http.Request) {

	// Подключаемся к базе данных
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Добавляем пользователя в базу данных
	_, err = db.Query("INSERT INTO users (name, password) VALUES ($1, $2)", "test", "test")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Редирект на страницу логина
	http.Redirect(w, r, "/login", http.StatusFound)
}	
