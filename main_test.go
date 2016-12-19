package goclean

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goclean/infrastructure/jwtauth"
	"goclean/interfaceadapter/controller"
	mdw "goclean/interfaceadapter/middleware"
	"goclean/interfaceadapter/repository"
	"goclean/usecase"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Sample integration to test the whole code
func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip Main Integration test")
	}

	// re-create system
	// Create repositories
	userRepo := repository.NewUserRepo()
	authRepo := repository.NewAuthRepo()

	// Create use case
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Create infrastructure
	jwtAuth := jwtauth.NewJwtAuth()

	// Create controller
	userCtrl := controller.NewUserCtrl(userUseCase)

	// Create middleware
	mdwChain := mdw.NewChain(mdw.MdwCORS, mdw.MdwLog, mdw.MdwHeader)
	mdwToken := mdw.NewMdwToken(authRepo, jwtAuth)

	// Register routes
	router := mux.NewRouter()
	router.Path("/users/{userId}").Methods("GET").Handler(
		mdwChain.Then(mdwToken.HandleFunc(userCtrl.GetUser)),
	)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected response code %v but got %v", 200, w.Code)
	}

	respBody, _ := ioutil.ReadAll(w.Body)
	respData := map[string]string{}
	_ = json.Unmarshal(respBody, &respData)

	if respData["id"] != "1" {
		t.Errorf("Expeccted id %v but got %v", "1", respData["id"])
	}
}
