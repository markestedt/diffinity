package main

import (
	"net/http"

	"github.com/markestedt/diffinity/internal/compare"
)

type compareViewModel struct {
	Result string
}

func (app *application) getIndexHandler(w http.ResponseWriter, r *http.Request) {
	app.templates.ExecuteTemplate(w, "pages/index", nil)
}

func (app *application) postCompareHandler(w http.ResponseWriter, r *http.Request) {
	input := compare.ParseRequest(r)

	html := compare.DiffHtml(input.OldApiSpec, input.NewApiSpec)
	viewModel := compareViewModel{
		Result: string(html[:]),
	}

	app.templates.ExecuteTemplate(w, "components/compareResult", viewModel)

}
