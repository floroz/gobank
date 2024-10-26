package models

type Account struct {
	ID            int     `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	AccountNumber int64   `json:"account_number"`
	Balance       float64 `json:"balance"`
	PhoneNumber   string  `json:"phone_number"`
	Country       string  `json:"country"`
}

type CreateAccountDTO struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	AccountNumber int64   `json:"account_number"`
	Balance       float64 `json:"balance"`
	PhoneNumber   string  `json:"phone_number"`
	Country       string  `json:"country"`
}

func NewCreateAccountDTO(firstName, lastName, phoneNumber, country string) *CreateAccountDTO {
	return &CreateAccountDTO{FirstName: firstName, LastName: lastName, Balance: 0, PhoneNumber: phoneNumber, Country: country}
}
