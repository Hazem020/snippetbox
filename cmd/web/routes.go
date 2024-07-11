package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/", http.StripPrefix("/static", fileServer))
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodGet, "/snippet/latest", app.snippetLatest)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreateForm)
	return app.recoverPanic(app.logRequest(secureHeaders(router)))
}
