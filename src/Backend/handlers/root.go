package handlers

import (
	"gochat/session"
	"net/http"
)

// Корневая страница
func HandleRoot(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {

			session, _ := session.Store.Get(r, session.SessionName)

			// Redirect to contacts if user is authenticated
			if session.Values["authenticated"] == true {
				http.Redirect(w, r, "/contacts/", http.StatusFound)
				return
			}

			// Redirect to login if user is not authenticated
			if session.Values["authenticated"] == false || session.Values["authenticated"] == nil {
				http.Redirect(w, r, "/login/", http.StatusFound)
				return
			}
		}
		// Serve the static files for other routes
		fs.ServeHTTP(w, r)
	}
}
