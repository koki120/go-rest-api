package main

import (
	"net/http"
	"os"

	"github.com/koki120/go-rest-api/adapter/database"
	"github.com/koki120/go-rest-api/api/router"
	"github.com/koki120/go-rest-api/interactor"
	"github.com/koki120/go-rest-api/log"
)

func main() {
	logger := log.NewLogger()

	db, err := database.NewMySQLDB()
	if err != nil {
		logger.Error("Failed to connect to database", err)
		os.Exit(1)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			logger.Error("Failed to close database connection", err)
		}
	}()

	memoUC := interactor.NewMemoUseCase()

	s := router.NewServer(memoUC)

	if err := http.ListenAndServe(":80", s); err != nil {
		logger.Error("Failed to start server", err)
		os.Exit(1)
	}

}
