package api

import (
	"log"
	"net/http"

	"github.com/floroz/gobank/api/account"
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
	mux := http.NewServeMux()
	accountRouter := account.NewAccountRouter()

	mux.HandleFunc("GET /account/{id}", accountRouter.GetAccount)
	mux.HandleFunc("POST /account", accountRouter.CreateAccount)

	log.Fatal(http.ListenAndServe(s.listenAddress, mux))
}
