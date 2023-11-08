package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/koki120/go-rest-api/log"
)

func Hello(w http.ResponseWriter, _ *http.Request) {
	logger := log.NewLogger()
	if err := json.NewEncoder(w).Encode(
		map[string]string{
			"message": "hello world!",
		},
	); err != nil {
		logger.Error("failed to encode json", err)
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
	}

}
