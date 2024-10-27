package storage

import (
	"database/sql"
	"log"

	"github.com/floroz/gobank/api/models"
)

type PostgreSQLStorage struct {
	DB *sql.DB
}

func (s *PostgreSQLStorage) CreateAccount(c *models.CreateAccountDTO) (*models.Account, error) {
	// TODO implement "countries" table so we can store the country_id in this table and join when needed to avoid string duplication

	var account models.Account
	err := s.DB.QueryRow(`
			INSERT INTO accounts (first_name, last_name, phone_number, country_code) 
			VALUES ($1, $2, $3, $4) 
			RETURNING id, first_name, last_name, phone_number, country_code, created_at`, c.FirstName, c.LastName, c.PhoneNumber, c.Country).Scan(&account)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &account, nil
}

// func (s *PostgreSQLStorage) DeleteAccount(id string) (*models.Account, error) {
// 	return &models.Account{}
// }
// func (s *PostgreSQLStorage) UpdateAccount(a *models.Account) (*models.Account, error) {
// 	return &models.Account{}
// }
// func (s *PostgreSQLStorage) GetAccountById(id string) (*models.Account, error) {
// 	rows, err := s.DB.Query("SELECT * FROM accounts WHERE id = $1", id)

// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	return &models.Account{}
// }

func NewPostgreSQLStorage() *PostgreSQLStorage {
	return &PostgreSQLStorage{}
}
