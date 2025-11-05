package initiator

import (
	"users/pkgs/logger"
)

func InitLogger() logger.Logger {
	return logger.NewLogger()
}
