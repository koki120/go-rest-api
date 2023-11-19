package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func GetQuery(r http.Request) url.Values {
	return r.URL.Query()
}

// Expecting path parameters to be placed at the end
func GetLastPathParameter(r http.Request) string {
	prm := strings.Split(r.URL.Path, "/")
	return prm[len(prm)-1]
}

func WriteStatusText(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, http.StatusText(statusCode))
}

func WriteJSONResponse(w http.ResponseWriter, response interface{}, statusCode int) error {
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)
}
