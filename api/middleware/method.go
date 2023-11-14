package middleware

import "net/http"

type HandleHTTPMethod struct {
	mux             *http.ServeMux
	handlerWrappers []HandlerWrapper
}

type HandlerWrapper func(http.HandlerFunc) http.HandlerFunc

func NewHandleHTTPMethod(mux *http.ServeMux, handlerWrappers []HandlerWrapper) *HandleHTTPMethod {
	return &HandleHTTPMethod{
		mux,
		handlerWrappers,
	}
}

const methodNotAllowed = "Method not allowed."

func (h *HandleHTTPMethod) GET(url string, handler http.HandlerFunc) {
	allowMethod(h, url, http.MethodGet, handler)
}

func (h *HandleHTTPMethod) POST(url string, handler http.HandlerFunc) {
	allowMethod(h, url, http.MethodPost, handler)
}

func (h *HandleHTTPMethod) PUT(url string, handler http.HandlerFunc) {
	allowMethod(h, url, http.MethodPut, handler)
}

func (h *HandleHTTPMethod) DELETE(url string, handler http.HandlerFunc) {
	allowMethod(h, url, http.MethodDelete, handler)
}

func (h *HandleHTTPMethod) HEAD(url string, handler http.HandlerFunc) {
	allowMethod(h, url, http.MethodHead, handler)
}

func (h *HandleHTTPMethod) OPTIONS(url string, handler http.HandlerFunc) {
	allowMethod(h, url, http.MethodOptions, handler)
}

func (h *HandleHTTPMethod) PATCH(url string, handler http.HandlerFunc) {
	allowMethod(h, url, http.MethodPatch, handler)
}

func (h *HandleHTTPMethod) TRACE(url string, handler http.HandlerFunc) {
	allowMethod(h, url, http.MethodTrace, handler)
}

func (h *HandleHTTPMethod) CONNECT(url string, handler http.HandlerFunc) {
	allowMethod(h, url, http.MethodConnect, handler)
}

func allowMethod(h *HandleHTTPMethod, url string, method string, f http.HandlerFunc) {
	hw := applyWrappers(f, h.handlerWrappers...)

	h.mux.HandleFunc(
		url,
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == method {
				hw(w, r)
				return
			}
			http.Error(w, methodNotAllowed, http.StatusMethodNotAllowed)
		},
	)
}

func applyWrappers(handler http.HandlerFunc, wrappers ...HandlerWrapper) http.HandlerFunc {
	for _, wrapper := range wrappers {
		handler = wrapper(handler)
	}
	return handler
}
