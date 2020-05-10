package http

import (
	"github.com/gorilla/mux"
	"github.com/lakkinzimusic/horse_maze/auth"
)

//RegisterHTTPEndpoints func
func RegisterHTTPEndpoints(router mux.Router, useCase auth.UseCase) {
	handler := NewHandler(useCase)
	router.HandleFunc("/auth/sign-up", handler.SignUp)
	router.HandleFunc("/sign-in", handler.SignIn)
}
