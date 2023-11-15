package router

import "net/http"

// This function invokes middleware and allows a specific HTTP method.
func handleMethod(h *HandleHTTPMethod, url string, methodRoutes []MethodRoute) {
	h.mux.HandleFunc(
		url,
		func(w http.ResponseWriter, r *http.Request) {
			if next := findHandler(methodRoutes, r.Method); next != nil {
				hw := applyWrappers(next, h.handlerWrappers...)
				hw.ServeHTTP(w, r)
				return
			}
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
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

func findHandler(methodRoutes []MethodRoute, method string) http.HandlerFunc {
	for _, route := range methodRoutes {
		if string(route.httpMethod) == method {
			return route.handlerFunc
		}
	}
	return nil
}
