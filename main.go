package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/floroz/gobank/api"
	"github.com/floroz/gobank/storage"
	"github.com/joho/godotenv"
)

func genDbConnectionString() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	return dsn
}

func loadEnv() {
	godotenv.Load()
}

func connectDb() {
	db, err := sql.Open("postgres", genDbConnectionString())

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		// do something here
	}
}

func main() {
	loadEnv()
	connectDb()
	storage := storage.NewPostgreSQLStorage()

	hostname := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	server := api.NewAPIServer(hostname, storage)
	server.Run()
}
