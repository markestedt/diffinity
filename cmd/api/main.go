package main

import (
	"embed"
	"errors"
	"log"
	"net/http"
	"text/template"
)

//go:embed templates
var templates embed.FS

type application struct {
	templates *template.Template
}

func main() {
	t := template.Must(template.ParseFS(templates, "templates/components/*.html"))
	t = template.Must(t.ParseFS(templates, "templates/pages/*.html"))

	app := &application{
		templates: t,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.getIndexHandler)
	mux.HandleFunc("POST /compare", app.postCompareHandler)

	err := http.ListenAndServe(":9494", mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server closed\n")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
