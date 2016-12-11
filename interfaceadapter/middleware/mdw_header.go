package middleware

import "net/http"

// Add Header to all response
func MdwHeader(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, private, max-age=0")
		handler.ServeHTTP(w, r)
	})
}
