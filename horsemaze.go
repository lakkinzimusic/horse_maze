package main

import (
	"log"
	"os"

	"github.com/lakkinzimusic/horse_maze/api/version"

	"github.com/lakkinzimusic/horse_maze/api/server"
)

func main() {

	log.Printf(
		"Starting the service on port %s...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	app := server.NewApp()
	if err := app.Run(port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
