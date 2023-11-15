package util

import (
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
