package goclean

import (
	"github.com/gorilla/mux"
	"goclean/interfaceadapter/controller"
	"goclean/interfaceadapter/repository"
	"goclean/usecase"
	"net/http"
)

func main() {

	// create reposilories
	userRepo := repository.NewUserRepo()

	// Create use case
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Create controller
	userCtrl := controller.NewUserCtrl(userUseCase)
	ctrl := controller.NewController(userCtrl)

	r := mux.NewRouter()
	ctrl.Register(r)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
