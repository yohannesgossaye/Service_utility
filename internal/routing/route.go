package routing

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
	Mid     []func(http.Handler) http.Handler
}

func NewRoute(r chi.Router, routes []Route) {
	for _, route := range routes {
		handler := http.Handler(route.Handler)
		for _, mid := range route.Mid {
			handler = mid(handler)
		}
		r.Method(route.Method, route.Path, handler)
	}
}
