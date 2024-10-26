package main

import (
	"github.com/floroz/gobank/api"
)

func main() {
	server := api.NewAPIServer("localhost:8080")
	server.Run()
}
