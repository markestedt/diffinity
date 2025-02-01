package compare

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/tufin/oasdiff/diff"
	"github.com/tufin/oasdiff/formatters"
)

func DiffHtml(s1 string, s2 string) []byte {
	loader := openapi3.NewLoader()

	spec1, _ := loader.LoadFromData([]byte(s1))
	spec2, _ := loader.LoadFromData([]byte(s2))

	config := diff.NewConfig().WithExcludeElements([]string{"examples"})

	d, _ := diff.Get(config, spec1, spec2)

	formatter := formatters.HTMLFormatter{}

	html, _ := formatter.RenderDiff(d, formatters.NewRenderOpts())
	return html
}
