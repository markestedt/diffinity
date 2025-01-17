package main

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/markestedt/openapidiff/internal/compare"
	"github.com/tufin/oasdiff/diff"
	"github.com/tufin/oasdiff/formatters"
)

type compareViewModel struct {
	Result string
}

func (app *application) getIndexHandler(w http.ResponseWriter, r *http.Request) {
	app.templates.ExecuteTemplate(w, "pages/index", nil)
}

func (app *application) postCompareHandler(w http.ResponseWriter, r *http.Request) {
	input := compare.ParseRequest(r)

	loader := openapi3.NewLoader()

	s1, _ := loader.LoadFromData([]byte(input.OldApiSpec))
	s2, _ := loader.LoadFromData([]byte(input.NewApiSpec))

	config := diff.NewConfig().WithExcludeElements([]string{"examples"})

	d, _ := diff.Get(config, s1, s2)

	formatter := formatters.HTMLFormatter{}

	html, _ := formatter.RenderDiff(d, formatters.NewRenderOpts())

	viewModel := compareViewModel{
		Result: string(html[:]),
	}

	app.templates.ExecuteTemplate(w, "components/compareResult", viewModel)

}
