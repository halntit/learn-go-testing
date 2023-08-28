package main

import (
	"flag"
	"log"
	"encoding/gob" // go binary encoding
	"net/http"
	"webapp/pkg/db"
	"webapp/pkg/data"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	DSN     string
	DB      db.PostgresConn
	Session *scs.SessionManager
}

func main() {
	gob.Register(data.User{})

	// set up an app config
	app := application{}

	flag.StringVar(&app.DSN, "dsn", "host=localhost user=postgres password=postgres dbname=users port=5432 sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // don't close until main exit

	app.DB = db.PostgresConn{DB: conn}

	// get a session manager
	app.Session = getSession()

	// print out message
	log.Println("Starting application... on port 8080")

	// start the application
	err = http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
