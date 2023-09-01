package domain

type Contact struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type ContactUseCase interface {
	GetContacts() ([]Contact, error)
}
