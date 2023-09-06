package main

import (
	"flag"
	"log"
	"fmt"
	"net/http"
	"webapp/pkg/repository"
	"webapp/pkg/repository/dbrepo"
)

const port = 8090

type application struct {
	DSN string
	DB repository.DatabaseRepo
	Domain string
	JWTSecret string
}

func main() {
	var app application
	flag.StringVar(&app.Domain, "domain", "example.com", "Domain foro app, e.g. company.com")
	flag.StringVar(&app.DSN, "dsn", "host=localhost user=postgres password=postgres dbname=users port=5432 sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "s3c4et", "signing secret")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // don't close until main exit

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	log.Printf("Starting api... on port %d\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
