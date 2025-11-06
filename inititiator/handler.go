package initiator

import (
	"users/internal/handler/http"
	"users/internal/handler/http/bills"
	"users/internal/handler/http/users"
	"users/pkgs/logger"
)

type Handle struct {
	UsersHandler http.UserHandler
	BillsHandler http.BillsHandler
	logger       logger.Logger
}

func NewHandler(service ServiceUs, log logger.Logger) Handle {
	return Handle{
		UsersHandler: users.NewUserH(
			service.UsersService,
			log,
		),
		BillsHandler: bills.NewBillH(
			service.BillService,
			log,
		),
	}
}
