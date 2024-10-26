package account

type Account struct {
	ID            int     `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	AccountNumber int64   `json:"account_number"`
	Balance       float64 `json:"balance"`
	// PhoneNumber   string
	// Country       string
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{FirstName: firstName, LastName: lastName}
}
