package handlers

import (
	"net/http"
)

// Корневая страница
func HandleRoot(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		// Serve the static files for other routes
		fs.ServeHTTP(w, r)
	}
}