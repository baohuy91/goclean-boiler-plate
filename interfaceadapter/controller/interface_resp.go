package controller

import "net/http"

type Response interface{
	Ok(w http.ResponseWriter, m interface{})
	Error(w http.ResponseWriter, statusCode int, err error)
}
