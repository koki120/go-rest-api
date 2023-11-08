package main

import (
	"net/http"
	"os"

	"github.com/koki120/go-rest-api/api/router"
	"github.com/koki120/go-rest-api/log"
)

func main() {
	logger := log.NewLogger()

	r := router.NewRouter()

	if err := http.ListenAndServe(":80", r); err != nil {
		logger.Error("Failed to start server", err)
		os.Exit(1)
	}

}
