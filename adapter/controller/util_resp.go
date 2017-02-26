package controller

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"net/http"
)

func ResponseOk(w http.ResponseWriter, m interface{}) {
	// TODO: wrap this to API response format
	js, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func ResponseError(w http.ResponseWriter, statusCode int, err error) {
	if statusCode == http.StatusInternalServerError {
		logrus.Error(err)
	} else {
		logrus.Debug(err)
	}
	w.WriteHeader(statusCode)
}
