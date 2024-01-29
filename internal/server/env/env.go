package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_ENGINE string
	DB_NAME   string
}

func LoadEnv() Env {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return Env{
		DB_ENGINE: os.Getenv("DB_ENGINE"),
		DB_NAME:   os.Getenv("DB_NAME"),
	}
}
