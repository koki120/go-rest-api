package middleware

import "net/http"

type (
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

func (h *HandleHTTPMethod) GET(url string, handler http.HandlerFunc) {
	handleMethod(h, url, http.MethodGet, handler)
}

func (h *HandleHTTPMethod) POST(url string, handler http.HandlerFunc) {
	handleMethod(h, url, http.MethodPost, handler)
}

func (h *HandleHTTPMethod) PUT(url string, handler http.HandlerFunc) {
	handleMethod(h, url, http.MethodPut, handler)
}

func (h *HandleHTTPMethod) DELETE(url string, handler http.HandlerFunc) {
	handleMethod(h, url, http.MethodDelete, handler)
}

func (h *HandleHTTPMethod) HEAD(url string, handler http.HandlerFunc) {
	handleMethod(h, url, http.MethodHead, handler)
}

func (h *HandleHTTPMethod) OPTIONS(url string, handler http.HandlerFunc) {
	handleMethod(h, url, http.MethodOptions, handler)
}

func (h *HandleHTTPMethod) PATCH(url string, handler http.HandlerFunc) {
	handleMethod(h, url, http.MethodPatch, handler)
}

func (h *HandleHTTPMethod) TRACE(url string, handler http.HandlerFunc) {
	handleMethod(h, url, http.MethodTrace, handler)
}

func (h *HandleHTTPMethod) CONNECT(url string, handler http.HandlerFunc) {
	handleMethod(h, url, http.MethodConnect, handler)
}
