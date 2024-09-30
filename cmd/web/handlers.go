package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	templatePaths := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(templatePaths...)

	if err != nil {
		app.serverError(w, r, err)
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) viewSnippet(w http.ResponseWriter, r *http.Request) {
	snippetId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || snippetId < 1 {

		app.logger.Error("invalid id", "id", r.PathValue("id"), "method", r.Method, "path", r.URL.RequestURI())
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display information about snippet with ID: %d", snippetId)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Display form to create a new snippet...")
}

func (app *application) createSnippetPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Create a new snippet...")
}
