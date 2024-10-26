package main

import (
	"github.com/floroz/gobank/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	server := api.NewAPIServer("localhost:8080")
	server.Run()
}
