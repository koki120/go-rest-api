package middleware

import (
	"net/http"

	"github.com/koki120/go-rest-api/log"
)

func Authentication(next http.Handler) http.HandlerFunc {
	logger := log.NewLogger()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("TODO:authentication")

		next.ServeHTTP(w, r)
	})
}
