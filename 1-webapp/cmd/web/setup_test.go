package main

import (
	"os"
	"log"
	"webapp/pkg/db"
	"testing"
)

var app application

func TestMain(m *testing.M) { // m for main
	pathToTemplates = "./../../templates"

	app.Session = getSession()

	app.DSN = "host=localhost user=postgres password=postgres dbname=users port=5432 sslmode=disable timezone=UTC connect_timeout=5"

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // don't close until main exit

	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run()) // run this first then other tests
}
