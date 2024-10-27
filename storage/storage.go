package storage

import (
	"github.com/floroz/gobank/api/models"
)

type Storager interface {
	CreateAccount(c *models.CreateAccountDTO) (*models.Account, error)
	// DeleteAccount(id string) (*models.Account, error)
	// UpdateAccount(a *models.Account) (*models.Account, error)
	// GetAccountById(id string) (*models.Account, error)
}
