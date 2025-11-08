package initiator

import (
	"users/internal/service"
	appbill "users/internal/service/bills"
	grpcClient "users/internal/service/bills/grpc_client"
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

	// Initialize gRPC client for bill payments
	grpcBillClient := grpcClient.NewGRPCBillClient()
	// Pass transaction repository to bills service
	billService := appbill.NewBills(p.transactions, grpcBillClient, log)

	return ServiceUs{
		UsersService: userService,
		BillService:  billService,
	}
}
