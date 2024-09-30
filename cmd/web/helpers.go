package main

import "net/http"

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err.Error(), "method", r.Method, "path", r.URL.RequestURI())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
