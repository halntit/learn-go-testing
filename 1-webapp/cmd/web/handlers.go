package main

import (
	"html/template"
	"net/http"
	"path"
	"log"
	"fmt"
)

var pathToTemplates = "./templates"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

func (app *application) About(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "", nil)
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println(email, password)
	fmt.Fprintf(w, "email: %s, password: %s", email, password)
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, tmpl string, data *TemplateData) error {
	// parse the template from disk
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, tmpl))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	data.IP = app.ipFromContext(r.Context())

	// execute the template, passing data
	// w.Header().Set("Content-Type", "text/html")
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
