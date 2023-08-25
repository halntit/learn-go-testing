package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	// set up an app config
	app := application{}

	// get a session manager
	app.Session = getSession()

	// print out message
	log.Println("Starting application... on port 8080")

	// start the application
	err := http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
