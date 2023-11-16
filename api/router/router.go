package router

import (
	"net/http"

	"github.com/koki120/go-rest-api/api/middleware"
	"github.com/koki120/go-rest-api/features/health"
	"github.com/koki120/go-rest-api/features/memo"
	"github.com/koki120/go-rest-api/interface/i_health"
	"github.com/koki120/go-rest-api/interface/i_memo"
)

func NewServer(
	healthUC i_health.UseCase,
	memoUC i_memo.UseCase,
) *http.ServeMux {
	mux := http.NewServeMux()
	healthHandler := health.NewHandler(healthUC)
	mux.HandleFunc("/health", healthHandler.Health)

	h := NewHandleHTTPMethod(mux, []HandlerWrapper{
		middleware.Authentication,
		middleware.Recovery,
	})

	memoHandler := memo.NewHandler(memoUC)

	h.Add("/memo/", []MethodRoute{
		{httpMethod: GET, handlerFunc: memoHandler.FindByID},
		{httpMethod: POST, handlerFunc: memoHandler.Create},
		{httpMethod: DELETE, handlerFunc: memoHandler.Delete},
		{httpMethod: PATCH, handlerFunc: memoHandler.Update},
	})
	h.Add("/memo", []MethodRoute{
		{httpMethod: GET, handlerFunc: memoHandler.Search},
	})

	return mux
}
