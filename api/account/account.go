package account

import (
	"fmt"
	"net/http"
)

type AccountRouter struct{}

func (s *AccountRouter) HandleAccount(w http.ResponseWriter, r *http.Request) {
}

func (s *AccountRouter) GetAccount(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("GET Account: ", id)
}

func (s *AccountRouter) CreateAccount(w http.ResponseWriter, r *http.Request) {
}

func (s *AccountRouter) DeleteAccount(w http.ResponseWriter, r *http.Request) {
}

func (s *AccountRouter) HandleTransfer(w http.ResponseWriter, r *http.Request) {
}
