package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"snippetbox.nathan-r-nicholson.com/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	_, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

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

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func (app *application) viewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	snippetId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || snippetId < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(snippetId)

	if err != nil {
		if err == models.ErrNoRecord {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/view.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	templateData := &templateData{Snippet: snippet}

	err = ts.ExecuteTemplate(w, "base", templateData)

	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Display form to create a new snippet...")
}

func (app *application) createSnippetPost(w http.ResponseWriter, r *http.Request) {

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
