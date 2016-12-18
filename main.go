package goclean

import (
	"github.com/gorilla/mux"
	"goclean/infrastructure"
	"goclean/infrastructure/jwtauth"
	"goclean/infrastructure/sendgridmail"
	"goclean/interfaceadapter/controller"
	mdw "goclean/interfaceadapter/middleware"
	"goclean/interfaceadapter/repository"
	"goclean/usecase"
	"net/http"
)

func main() {

	// Create repositories
	userRepo := repository.NewUserRepo()
	authRepo := repository.NewAuthRepo()

	// Create use case
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Create infrastructure Api response
	response := infrastructure.ApiResponse{}
	jwtAuth := jwtauth.NewJwtAuth()
	// Get these info from config file and add here
	mailManager := sendgridmail.NewSendGridMailManager("host", "endpoint", "apikey")

	// Create controller
	userCtrl := controller.NewUserCtrl(response, userUseCase)
	authCtrl := controller.NewAuthCtrl(response, userUseCase, authRepo, jwtAuth, mailManager)

	// Create middle ware
	mdwChain := mdw.NewChain(mdw.MdwCORS, mdw.MdwLog, mdw.MdwHeader)
	mdwToken := mdw.NewMdwToken(response, authRepo, jwtAuth)

	// Register routes
	r := mux.NewRouter()
	r.Path("/auth/registerbyemail").Methods("POST").Handler(
		mdwChain.Then(http.HandlerFunc(authCtrl.RegisterByMail)),
	)
	r.Path("/users/{userId}").Methods("GET").Handler(
		mdwChain.Then(mdwToken.HandleFunc(userCtrl.GetUser)),
	)

	// Start handle request
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
