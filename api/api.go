package api

import (
	"log"
	"net/http"

	"github.com/floroz/gobank/api/middlewares"
	"github.com/floroz/gobank/api/models"
	"github.com/floroz/gobank/api/resources/account"
)

type PostgreSQLStorage struct{}

func (s *PostgreSQLStorage) CreateAccount(c *models.CreateAccountDTO) (*models.Account, error) {
	return &models.Account{}
}
func (s *PostgreSQLStorage) DeleteAccount(id string) (*models.Account, error) {
	return &models.Account{}
}
func (s *PostgreSQLStorage) UpdateAccount(a *models.Account) (*models.Account, error) {
	return &models.Account{}
}
func (s *PostgreSQLStorage) GetAccountById(id string) (*models.Account, error) {
	return &models.Account{}
}

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
	storage := PostgreSQLStorage{}
	accountRouter := account.NewAccountRouter(&storage)

	mux.HandleFunc("GET /account/{id}", accountRouter.GetAccount)
	mux.HandleFunc("POST /account", accountRouter.CreateAccount)

	loggedMux := middlewares.LoggerMiddleware(mux)

	log.Println("Server running on: ", s.listenAddress)
	log.Fatal(http.ListenAndServe(s.listenAddress, loggedMux))
}
