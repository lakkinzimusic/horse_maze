package handlers

import (
	"github.com/gorilla/mux"
)

// Router register necessary routes and returns an instance of a router.
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Base)
	r.HandleFunc("/home/{letter}/{number:[0-9]+}", Home).Methods("GET")
	return r
}
