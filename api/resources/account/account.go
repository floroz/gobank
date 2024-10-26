package account

type Account struct {
	ID            int
	FirstName     string
	LastName      string
	AccountNumber int64
	Balance       float64
	// PhoneNumber   string
	// Country       string
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{FirstName: firstName, LastName: lastName}
}
