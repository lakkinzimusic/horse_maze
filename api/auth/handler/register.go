package http

import (
	"github.com/gorilla/mux"
	"github.com/lakkinzimusic/horse_maze/api/auth"
)

//RegisterHTTPEndpoints func
func RegisterHTTPEndpoints(router *mux.Router, useCase auth.UseCase) {
	handler := NewHandler(useCase)
	router.HandleFunc("/sign-up", handler.SignUp).Methods("POST")
	router.HandleFunc("/sign-in", handler.SignIn)
}
