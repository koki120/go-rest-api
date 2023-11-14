package middleware

import "net/http"

const methodNotAllowed = "Method not allowed."

// This function invokes middleware and allows a specific HTTP method.
func allowMethod(h *HandleHTTPMethod, url string, method string, next http.HandlerFunc) {
	h.mux.HandleFunc(
		url,
		func(w http.ResponseWriter, r *http.Request) {
			hw := applyWrappers(next, h.handlerWrappers...)
			if r.Method == method {
				hw.ServeHTTP(w, r)
				return
			}
			http.Error(w, methodNotAllowed, http.StatusMethodNotAllowed)
		},
	)
}

func applyWrappers(next http.HandlerFunc, wrappers ...HandlerWrapper) http.Handler {
	for _, wrapper := range wrappers {
		next = wrapper(next)
	}
	return next
}
