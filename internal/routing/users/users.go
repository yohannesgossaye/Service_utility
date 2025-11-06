package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	userhandler "users/internal/handler/http"
	usersrouting "users/internal/routing"
)

func InitRoutes(router chi.Router, handler userhandler.UserHandler) {
	router.Route("/users", func(r chi.Router) {
		routes := []usersrouting.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: handler.CreateUser,
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/{id}",
				Handler: handler.GetUser,
			},
			{
				Method:  http.MethodGet,
				Path:    "/users",
				Handler: handler.GetUsers,
			},
			{
				Method:  http.MethodPut,
				Path:    "/update/{id}",
				Handler: handler.UpdateUser,
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete/{id}",
				Handler: handler.DeleteUser,
			},
			{
				Method:  http.MethodPatch,
				Path:    "/enable/{id}",
				Handler: handler.EnableAccount,
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: handler.LoginUser,
			},
		}
		usersrouting.NewRoute(r, routes)
	})

}
