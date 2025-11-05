package initiator

import (
	"users/internal/service"
	appUsers "users/internal/service/users"
	"users/pkgs/logger"
	"users/pkgs/utils/email"
)

type ServiceUs struct {
	UsersService service.UserS
}

func InitService(p Persistence, sender email.Sender, log logger.Logger) ServiceUs {

	userService := appUsers.InitSvc(p.users, log, sender)

	return ServiceUs{
		UsersService: userService,
	}
}
