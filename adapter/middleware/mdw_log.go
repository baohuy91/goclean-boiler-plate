package middleware

import (
	"context"
	"crypto/rand"
	"fmt"
	"goclean/adapter"
	"net/http"
)

// TODO: Working in progress
type MdwLog interface {
	ChainFunc(h http.Handler) http.Handler
}

func NewMdwLog(logger adapter.Logger) MdwLog {
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
	logger adapter.Logger
}

func (m *mdwLogImpl) ChainFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request with request Id from client or generate a new one
		reqId := r.Header.Get("X-Request-ID")
		if reqId == "" {
			reqId = generateRandomID()
		}
		r.WithContext(context.WithValue(r.Context(), "requestId", reqId))

		// Log request
		//m.logger.LogWithFields(map[string]interface{}{
		//	"requestId":  reqId,
		//	"method":     r.Method,
		//	"requestUri": r.RequestURI,
		//}, "Request")

		//rW := &writerWrapper{-1, w}
		h.ServeHTTP(w, r)

		// Log response
		//m.logger.LogWithFields(map[string]interface{}{
		//	"requestId":  reqId,
		//	"httpStatus": rW.status,
		//}, "Reponse")
	})
}

func generateRandomID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
