package initiator

import (
	routing "users/internal/routing/users"

	"github.com/go-chi/chi/v5"
)

func InitHandlers(r chi.Router, handler Handle) {
	r.Route("/api/v1", func(r chi.Router) {

		routing.InitRoutes(r, handler.UsersHandler)
	})
}
