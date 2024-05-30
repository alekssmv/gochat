package middlewares

import (
	"gochat/session"
	"net/http"
)

// Проверка авторизации
func NotAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, _ := session.Store.Get(r, session.SessionName)

		// Редирект на /contacts если пользователь уже авторизован
		if session.Values["authenticated"] == true {
			http.Redirect(w, r, "/contacts", http.StatusFound)
			return
        }
        next.ServeHTTP(w, r)
    })
}
