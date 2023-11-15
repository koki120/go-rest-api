package router

import (
	"net/http"

	"github.com/koki120/go-rest-api/api/middleware"
	"github.com/koki120/go-rest-api/features/hello"
)

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()

	h := NewHandleHTTPMethod(mux, []HandlerWrapper{
		middleware.Authentication,
		middleware.Recovery,
	})

	helloHandler := hello.NewHandler()

	h.Add("/hello", []MethodRoute{
		{httpMethod: GET, handlerFunc: helloHandler.Hello},
		{httpMethod: DELETE, handlerFunc: helloHandler.Hello},
	})

	return mux
}
