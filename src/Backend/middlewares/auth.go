package middlewares

import (
	"net/http"
	"gochat/session"
)

// Проверка авторизации
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, _ := session.Store.Get(r, session.SessionName)

		// Редирект на / если пользователь не авторизован
		if session.Values["authenticated"] != true {
			http.Redirect(w, r, "/login/", http.StatusFound)
			return
        }
        next.ServeHTTP(w, r)
    })
}
