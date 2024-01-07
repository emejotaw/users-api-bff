package entity

type Address struct {
	ID           string
	State        string
	City         string
	Neighborhood string
	Street       string
	Number       string
	ZipCode      string
	UserID       string
	User         *User
}
