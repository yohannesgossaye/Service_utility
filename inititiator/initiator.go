package initiator

import (
	"fmt"
	"log"
	"net/http"

	"users/pkgs/logger"

	"users/pkgs/utils/email"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func Initiator() {

	_ = godotenv.Load()

	logger := logger.NewLogger()

	fmt.Println(" Initializing configuration...")
	cfg := InitConfig()
	logger.Infof(" Configuration initialized")

	logger.Infof(" Initializing persistence...")
	mongo, err := InitPersistence(cfg.MongoURI, cfg.Database, cfg.UsersColl)
	if err != nil {
		log.Fatalf("❌ Failed to initialize persistence: %v", err)
	}
	logger.Infof(" Persistence initialized")
	email := email.NewSMTPSender()
	logger.Infof(" Initializing service layer...")
	service := InitService(mongo, email, logger)
	logger.Infof(" Service layer initialized")

	logger.Infof(" Initializing handlers...")
	handler := NewHandler(service, logger)
	logger.Infof(" Handlers initialized")

	r := chi.NewRouter()
	InitHandlers(r, handler)
	logger.Infof(" Routes initialized")

	logger.Infof("🚀 Starting server on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r))
}
