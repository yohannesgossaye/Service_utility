package initiator

import (
	"os"
)

type Config struct {
	Port      string
	MongoURI  string
	Database  string
	UsersColl string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func InitConfig() Config {
	return Config{
		Port:      getEnv("PORT", "8080"),
		MongoURI:  getEnv("MONGO_URI", "mongodb://db_mongo:27017"),
		Database:  getEnv("MONGO_DB", "users_db"),
		UsersColl: getEnv("MONGO_COLLECTION", "user_col"),
	}
}
