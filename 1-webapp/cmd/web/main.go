package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	// set up an app config
	app := application{}

	// get application routes
	mux := app.routes()

	// print out message
	log.Println("Starting application... on port 8080")

	// start the application
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
