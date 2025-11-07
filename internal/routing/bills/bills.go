package bills

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	billshandler "users/internal/handler/http"
	middleware "users/internal/handler/middleware"
	billsrouting "users/internal/routing"
)

func Billroutes(router chi.Router, handler billshandler.BillsHandler, mid middleware.AuthMiddleware) {
	router.Route("/bills", func(r chi.Router) {
		routes := []billsrouting.Route{
			{
				Method:  http.MethodPost,
				Path:    "/findbill",
				Handler: handler.GetBills,
				Mid: []func(next http.Handler) http.Handler{
					mid.AuthenticateToken,
				},
			},
			{
				Method:  http.MethodPost,
				Path:    "/paybills",
				Handler: handler.PayBills,
				Mid: []func(next http.Handler) http.Handler{
					mid.AuthenticateToken,
				},
			},
		}
		billsrouting.NewRoute(r, routes)
	})
}
