package initiator

import (
	"users/internal/service"
	appbill "users/internal/service/bills"
	appUsers "users/internal/service/users"
	"users/pkgs/logger"
	"users/pkgs/utils/email"
)

type ServiceUs struct {
	UsersService service.UserS
	BillService  service.BillService
}

func InitService(p Persistence, sender email.Sender, log logger.Logger) ServiceUs {

	userService := appUsers.InitSvc(p.users, log, sender)
	billService := appbill.NewBills(log)

	return ServiceUs{
		UsersService: userService,
		BillService:  billService,
	}
}
