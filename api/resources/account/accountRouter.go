package account

import (
	"encoding/json"
	"net/http"
)

type accountRouter struct{}

func (s *accountRouter) HandleAccount(w http.ResponseWriter, r *http.Request) {
}

func (s *accountRouter) GetAccount(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OKAY")
}

func (s *accountRouter) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// get the information from the payload

	// add the query to the db

	// get the returned value from the  db

	account := NewAccount("Hey", "Joe")

	// return it to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(account)
}

func (s *accountRouter) DeleteAccount(w http.ResponseWriter, r *http.Request) {
}

func (s *accountRouter) HandleTransfer(w http.ResponseWriter, r *http.Request) {
}

func NewAccountRouter() *accountRouter {
	return &accountRouter{}
}
