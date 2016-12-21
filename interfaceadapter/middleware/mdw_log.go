package middleware

import (
	"goclean/interfaceadapter"
	"net/http"
)

type MdwLog interface {
	ChainFunc(h http.Handler) http.Handler
}

func NewMdwLog(logger interfaceadapter.Logger) MdwLog {
	return &mdwLogImpl{
		logger: logger,
	}
}

type writerWrapper struct {
	status int
	http.ResponseWriter
}

// Wrapper to get status code on response
func (w *writerWrapper) WriteHeader(statusCode int) {
	w.status = statusCode
	w.WriteHeader(statusCode)
}

type mdwLogImpl struct {
	logger interfaceadapter.Logger
}

func (m *mdwLogImpl) ChainFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request
		m.logger.LogWithFields(map[string]interface{}{
			"method":     r.Method,
			"requestUri": r.RequestURI,
		}, "Request")

		rW := &writerWrapper{-1, w}
		h.ServeHTTP(rW, r)

		// Log response
		m.logger.LogWithFields(map[string]interface{}{
			"httpStatus": rW.status,
			"requestUri": r.RequestURI,
		}, "Reponse")
	})
}
