package main

import (
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.nathan-r-nicholson.com/internal/models"
	"snippetbox.nathan-r-nicholson.com/internal/validator"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	templateData := newTemplateData(r)
	templateData.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl", templateData)
}

func (app *application) viewSnippet(w http.ResponseWriter, r *http.Request) {
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

	templateData := newTemplateData(r)
	templateData.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl", templateData)
}

type snippetCreateForm struct {
	Title   string
	Content string
	Expires int
	validator.Validator
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	templateData := newTemplateData(r)

	templateData.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, http.StatusOK, "create.tmpl", templateData)
}

func (app *application) createSnippetPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.logger.Error("failed to parse form", "error", err)
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))

	if err != nil {
		app.logger.Error("expire date cannot be parsed to number", "error", err)
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := snippetCreateForm{
		Title:   title,
		Content: content,
		Expires: expires,
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "Title is required")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "Title is too long")
	form.CheckField(validator.NotBlank(form.Content), "content", "Content is required")
	form.CheckField(validator.PermittedValues(form.Expires, 1, 7, 365), "expires", "Expiry must be 1, 7, or 365 days")

	if !form.Valid() {
		templateData := newTemplateData(r)
		templateData.Form = form
		app.logger.Error("form validation failed")
		app.render(w, r, http.StatusBadRequest, "create.tmpl", templateData)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
