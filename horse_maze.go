package main

import (
	"log"
	"my_projects/horse_maze/handlers"
	"my_projects/horse_maze/version"
	"net/http"
	// "rsc.io/quote"
)

func main() {

	log.Printf(
		"Starting the service on port %s...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	log.Fatal("Port is not set.")
	// }
	// fmt.Println(port)
	router := handlers.Router()
	log.Fatal(http.ListenAndServe(":"+"8000", router))
}
