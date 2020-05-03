package main

import (
	"log"
	"my_projects/horse_maze/handlers"
	"net/http"
	// "rsc.io/quote"
)

func main() {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	log.Fatal("Port is not set.")
	// }

	router := handlers.Router()
	log.Fatal(http.ListenAndServe(":"+"8080", router))
}
