package router

import "net/http"

type (
	MethodRoute struct {
		httpMethod  HTTPMethod
		handlerFunc http.HandlerFunc
	}

	HandleHTTPMethod struct {
		mux             *http.ServeMux
		handlerWrappers []HandlerWrapper
	}

	HandlerWrapper func(next http.Handler) http.HandlerFunc
)

func NewHandleHTTPMethod(mux *http.ServeMux, handlerWrappers []HandlerWrapper) *HandleHTTPMethod {
	return &HandleHTTPMethod{
		mux,
		handlerWrappers,
	}
}

func (h *HandleHTTPMethod) Add(url string, methodRoutes []MethodRoute) {
	handleMethod(h, url, methodRoutes)
}
