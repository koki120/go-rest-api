package middleware

import (
	"net/http"

	"github.com/koki120/go-rest-api/log"
)

func Recovery(next http.Handler) http.HandlerFunc {
	logger := log.NewLogger()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic recovered:", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
