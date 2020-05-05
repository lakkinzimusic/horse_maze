package handlers

import (
	"fmt"
	"net/http"
)

//Base func
func Base(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello world. It's learning game")
}
