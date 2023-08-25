package main

import (
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) { // m for main
	pathToTemplates = "./../../templates"

	app.Session = getSession()

	os.Exit(m.Run()) // run this first then other tests
}
