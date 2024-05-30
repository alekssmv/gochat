package middlewares

import (
	"net/http"
	"gochat/session"
	"fmt"
)

// Проверка авторизации
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, _ := session.Store.Get(r, session.SessionName)

		fmt.Println("Auth: ", session.Values["authenticated"])

		if session.Values["authenticated"] != true {
			http.Redirect(w, r, "/", http.StatusFound)
			return
        }
        next.ServeHTTP(w, r)
    })
}
