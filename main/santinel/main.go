package main

import (
	"log"
	"net/http"

	"github.com/livghit/santinel/internal/server"
)

func main() {
	srv := server.NewServer()

	log.Printf("Server up and running on port %s", ":3000")
	log.Fatal(http.ListenAndServe(":3000", srv))
}
