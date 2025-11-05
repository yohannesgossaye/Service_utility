package routing

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Route struct {
	Method  string
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func NewRoute(router chi.Router, routes []Route) {
	for _, route := range routes {
		router.Method(route.Method, route.Path, http.HandlerFunc(route.Handler))
	}
}
