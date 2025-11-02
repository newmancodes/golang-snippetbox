package main

import (
	"net/http"
)

// Update the signature for the routes() method so that it returns a
// http.Handler instead of *http.ServeMux.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Create a file server which files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes as normal.
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}/{$}", app.snippetView)
	mux.HandleFunc("GET /snippet/create/{$}", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create/{$}", app.snippetCreatePost)

	// Pass the servemux as the 'next' parameter to the commonHeaders middleware.
	// Because commonHeaders is just a function, and the function returns a
	// http.Handler we don't need to do anything else.
	return app.logRequest(commonHeaders(mux))
}
