package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"snippetbox.newman.digital/internal/models"
	"snippetbox.newman.digital/ui"
)

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time value.
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap value and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup table mapping names to
// functions.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// Define a templateData type to act as the holding structure for
// and dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the project progresses.
type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string // Add a flash field to the templateData struct.
	IsAuthenticated bool
	CSRFToken       string
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	// Return the map.
	return cache, nil
}
