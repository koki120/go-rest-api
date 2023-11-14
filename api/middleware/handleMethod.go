package middleware

import "net/http"

const methodNotAllowed = "Method not allowed."

// This function invokes middleware and allows a specific HTTP method.
func handleMethod(h *HandleHTTPMethod, url string, method string, next http.HandlerFunc) {
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
	// To execute array elements (functions) in reverse order
	for i := len(wrappers) - 1; i >= 0; i-- {
		next = wrappers[i](next)
	}
	return next
}
