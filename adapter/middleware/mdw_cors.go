package middleware

import "net/http"

type MdwCORS interface {
	ChainFunc(handler http.Handler) http.Handler
}

func NewMdwCORS() MdwCORS {
	return &mdwCORSImpl{}
}

type mdwCORSImpl struct{}

// Handle CORS request fow API for client
func (m *mdwCORSImpl) ChainFunc(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}
		// Lets Gorilla work
		handler.ServeHTTP(w, r)
	})
}
