package database

import (
	"database/sql"

	"github.com/livghit/santinel/internal/server/env"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	connection *sql.DB
	name       string
}

func New(env env.Env) *Database {
	var conn *sql.DB

	switch env.DB_ENGINE {
	case "mysql":
		panic("MYSQL NOT IMPLEMENTED")
	case "sqlite":
		c, err := sql.Open("sqlite3", env.DB_NAME)
		if err != nil {
			panic(err)
		}
		conn = c
	case "postgress":
		panic("POSTGRESS NOT IMPLEMENTED")

	default:
		panic("Wrong engine selected please look at youre .env file .")
	}

	return &Database{
		connection: conn,
		name:       env.DB_NAME,
	}
}
