package storage

import "github.com/floroz/gobank/api/models"

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

func NewPostgreSQLStorage() *PostgreSQLStorage {
	return &PostgreSQLStorage{}
}
