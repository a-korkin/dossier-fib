package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
