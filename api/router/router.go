package router

import (
	"net/http"

	"github.com/koki120/go-rest-api/api/middleware"
	"github.com/koki120/go-rest-api/features/memo"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	h := middleware.NewHandleHTTPMethod(mux, []middleware.HandlerWrapper{
		middleware.Authentication,
	})

	memoHandler := memo.NewHandler()

	h.GET(MEMO, memoHandler.FindByID)

	return mux
}
