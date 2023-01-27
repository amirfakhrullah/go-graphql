package env

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var PORT string
var DB_URL string

func getEnvSafely(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic("environment variable missing :: " + k)
	}
	return v
}

const defaultPort = "8080"

func InitEnv() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = defaultPort
	}

	DB_URL = getEnvSafely("DB_URL")
}
