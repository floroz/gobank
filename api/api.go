package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/floroz/gobank/api/middlewares"
	"github.com/floroz/gobank/api/resources/account"
	"github.com/floroz/gobank/storage"
	_ "github.com/lib/pq"
)

type APIServer struct {
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func (s *APIServer) Run() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		// do something here
	}

	mux := http.NewServeMux()
	storage := storage.NewPostgreSQLStorage()
	accountRouter := account.NewAccountRouter(storage)

	mux.HandleFunc("GET /account/{id}", accountRouter.GetAccount)
	mux.HandleFunc("POST /account", accountRouter.CreateAccount)

	loggedMux := middlewares.LoggerMiddleware(mux)

	log.Println("Server running on: ", s.listenAddress)
	log.Fatal(http.ListenAndServe(s.listenAddress, loggedMux))
}
