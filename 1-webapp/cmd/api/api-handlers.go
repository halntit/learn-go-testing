package main

import (
	"net/http"
)

// var pathToTemplates = "./templates"

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read a JSON payload

	// look up user by email address

	// check password

	// generate tokens

	// send tokens to client

}

func (app *application) refresh(w http.ResponseWriter, r *http.Request) {

}

func (app *application) allUsers(w http.ResponseWriter, r *http.Request) {

}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {

}

func (app *application) insertUser(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {

}