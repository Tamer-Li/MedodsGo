package envRead

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

// readEnv finds the env file and creates the structure AppConfig
func readEnv(env string) (*AppConfig, error) {
	if err := godotenv.Load(env); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &AppConfig{
		DBHost: os.Getenv("DBHost"),
		DBPort: os.Getenv("DBPort"),
		DBUser: os.Getenv("DBUser"),
		DBPass: os.Getenv("DBPass"),
		DBName: os.Getenv("DBName"),
	}, nil
}
