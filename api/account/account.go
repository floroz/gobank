package account

import (
	"fmt"
	"net/http"
)

type accountRouter struct{}

func (s *accountRouter) HandleAccount(w http.ResponseWriter, r *http.Request) {
}

func (s *accountRouter) GetAccount(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("GET Account: ", id)
}

func (s *accountRouter) CreateAccount(w http.ResponseWriter, r *http.Request) {
}

func (s *accountRouter) DeleteAccount(w http.ResponseWriter, r *http.Request) {
}

func (s *accountRouter) HandleTransfer(w http.ResponseWriter, r *http.Request) {
}

func NewAccountRouter() *accountRouter {
	return &accountRouter{}
}
