package entity

type User struct {
	ID         string
	Name       string
	Email      string
	DocumentId string
	Address    *Address
}
