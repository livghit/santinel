package main

import (
	"log"
	"net/http"

	"github.com/livghit/santinel/internal/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	srv := server.NewServer(db)

	log.Printf("Server up and running on port %s", ":3000")
	log.Fatal(http.ListenAndServe(":3000", srv))
}

func initDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panicf("Database connection failed with error : %s", err)
	}

	return db
}
