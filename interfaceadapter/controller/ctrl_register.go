package controller

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type RegisterCtrl interface {
	RegisterByMail(w http.ResponseWriter, r *http.Request)
}

type registerCtrlImpl struct {
	response Response
}

type registerByMailReq struct {
	email string `json:"email"`
	pass  string `json:"pass"`
}

func (c *registerCtrlImpl) RegisterByMail(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.response.Error(w, http.StatusBadRequest, err)
		return
	}
	req := registerByMailReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		c.response.Error(w, http.StatusBadRequest, err)
		return
	}

	// Validate data
	// TODO: validate email format

	// TODO: validate password strength


}

