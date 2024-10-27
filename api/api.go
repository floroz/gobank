package api

import (
	"log"
	"net/http"

	"github.com/floroz/gobank/api/middlewares"
	"github.com/floroz/gobank/api/resources/account"
	"github.com/floroz/gobank/storage"
	_ "github.com/lib/pq"
)

type APIServer struct {
	listenAddress string
	storage       storage.Storager
}

func NewAPIServer(listenAddress string, storage storage.Storager) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		storage:       storage,
	}
}

func (s *APIServer) Run() {
	mux := http.NewServeMux()
	accountRouter := account.NewAccountRouter(s.storage)

	mux.HandleFunc("GET /account/{id}", accountRouter.GetAccount)
	mux.HandleFunc("POST /account", accountRouter.CreateAccount)

	loggedMux := middlewares.LoggerMiddleware(mux)

	log.Println("Server running on: ", s.listenAddress)
	log.Fatal(http.ListenAndServe(s.listenAddress, loggedMux))
}
