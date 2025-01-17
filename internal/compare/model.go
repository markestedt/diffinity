package compare

import "net/http"

type CompareRequest struct {
	OldApiSpec string
	NewApiSpec string
}

func ParseRequest(r *http.Request) CompareRequest {
	r.ParseForm()

	return CompareRequest{
		OldApiSpec: r.FormValue("oldApiSpec"),
		NewApiSpec: r.FormValue("newApiSpec"),
	}
}
