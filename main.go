package main

import (
	"net/http"
	"os"

	"github.com/koki120/go-rest-api/api/router"
	"github.com/koki120/go-rest-api/interactor"
	"github.com/koki120/go-rest-api/log"
)

func main() {
	logger := log.NewLogger()

	memoUC := interactor.NewMemoUseCase()

	s := router.NewServer(memoUC)

	if err := http.ListenAndServe(":80", s); err != nil {
		logger.Error("Failed to start server", err)
		os.Exit(1)
	}

}
