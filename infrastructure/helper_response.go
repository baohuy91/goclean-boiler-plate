package infrastructure

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
}

func (a ApiResponse) Ok(w http.ResponseWriter, m interface{}) {
	// TODO: wrap this to API response format
	js, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (a ApiResponse) Error(w http.ResponseWriter, statusCode int, err error) {
	if statusCode == http.StatusInternalServerError {
		// Log error somewhere
	}
	w.WriteHeader(statusCode)
}
