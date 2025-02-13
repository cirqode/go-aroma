package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir(("./internal/ui/static")))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return app.logRequest(commonHeaders(mux))

}
