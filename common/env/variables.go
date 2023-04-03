package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_DBNAME   string
	DB_PORT     string
	DB_SSL      string

	SECRET_KEY string

	ROUTER string
)

func Load() {
	godotenv.Load(".env")

	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_DBNAME = os.Getenv("DB_DBNAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_SSL = os.Getenv("DB_SSL")

	SECRET_KEY = os.Getenv("SECRET_KEY")

	ROUTER = os.Getenv("ROUTER")
}
