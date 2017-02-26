package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"goclean/adapter/controller"
	mdw "goclean/adapter/middleware"
	"goclean/adapter/repository"
	"goclean/infrastructure/jwtauth"
	"goclean/infrastructure/sendgridmail"
	"goclean/usecase"
	"net/http"
)

func main() {

	// Create repositories
	userRepo := repository.NewUserRepo()
	authRepo := repository.NewAuthRepo()

	// Create use case
	userUseCase := usecase.NewUserUseCase(userRepo)

	jwtAuth := jwtauth.NewJwtAuth()
	// Get these info from config file and add here
	mailManager := sendgridmail.NewSendGridMailManager("host", "endpoint", "apikey")

	// Create controller
	userCtrl := controller.NewUserCtrl(userUseCase)
	authCtrl := controller.NewAuthCtrl(userUseCase, authRepo, jwtAuth, mailManager)

	// Create middle ware
	mdwHeader := mdw.NewMdwHeader()
	mdwCORS := mdw.NewMdwCORS()
	mdwChain := mdw.NewChain(mdwCORS.ChainFunc, mdwHeader.ChainFunc)
	mdwToken := mdw.NewMdwToken(authRepo, jwtAuth)

	// Register routes
	r := mux.NewRouter()
	r.Path("/auth/registerbyemail").Methods("POST").Handler(
		mdwChain.Then(http.HandlerFunc(authCtrl.RegisterByMail)),
	)
	r.Path("/auth/login").Methods("POST").Handler(
		mdwChain.Then(http.HandlerFunc(authCtrl.LoginByEmail)),
	)
	r.Path("/auth/reqresetpass").Methods("POST").Handler(
		mdwChain.Then(http.HandlerFunc(authCtrl.RequestResetPassword)),
	)
	r.Path("/auth/resetpass").Methods("POST").Handler(
		mdwChain.Then(http.HandlerFunc(authCtrl.ResetPassword)),
	)

	// Need authorization
	r.Path("/users/{userId}").Methods("GET").Handler(
		mdwChain.Then(mdwToken.HandleFunc(userCtrl.GetUser)),
	)

	fmt.Println("Listen & serve on localhost:8080")
	// Start handle request
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
