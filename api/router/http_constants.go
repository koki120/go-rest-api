package router

import "net/http"

type HTTPMethod string

const (
	GET     HTTPMethod = http.MethodGet
	POST    HTTPMethod = http.MethodPost
	PUT     HTTPMethod = http.MethodPut
	DELETE  HTTPMethod = http.MethodDelete
	HEAD    HTTPMethod = http.MethodHead
	PATCH   HTTPMethod = http.MethodPatch
	OPTIONS HTTPMethod = http.MethodOptions
	TRACE   HTTPMethod = http.MethodTrace
)
