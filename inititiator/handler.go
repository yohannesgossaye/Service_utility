package initiator

import (
	"users/internal/handler/http"
	"users/internal/handler/http/users"
)

type Handle struct {
	UsersHandler http.UserHandler
}

func NewHandler(service ServiceUs) Handle {
	return Handle{
		UsersHandler: users.NewUserH(service.UsersService),
	}
}
