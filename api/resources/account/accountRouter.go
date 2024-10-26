package account

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/floroz/gobank/api/models"
	"github.com/floroz/gobank/storage"
)

type accountRouter struct {
	storage storage.Storager
}

func (s *accountRouter) HandleAccount(w http.ResponseWriter, r *http.Request) {
}

func (s *accountRouter) GetAccount(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OKAY")
}

func (s *accountRouter) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var accountDto models.CreateAccountDTO

	err := json.NewDecoder(r.Body).Decode(&accountDto)

	log.Println("account dto: ", &accountDto)

	if err != nil {
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request payload"})
		return
	}

	account, err := s.storage.CreateAccount(&accountDto)

	if err != nil {
		log.Println("Error creating the account", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "could not create account"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func (s *accountRouter) DeleteAccount(w http.ResponseWriter, r *http.Request) {
}

func (s *accountRouter) HandleTransfer(w http.ResponseWriter, r *http.Request) {
}

func NewAccountRouter(storage storage.Storager) *accountRouter {
	return &accountRouter{
		storage: storage,
	}
}
