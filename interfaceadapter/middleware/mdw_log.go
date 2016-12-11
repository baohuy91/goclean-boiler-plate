package middleware

import "net/http"

func MdwLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: add code here
		// logrus.Info(r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	});
}