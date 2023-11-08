package router

import (
	"net/http"

	"github.com/koki120/go-rest-api/api/handler"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handler.Hello)

	return mux
}
