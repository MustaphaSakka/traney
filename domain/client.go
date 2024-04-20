package domain

type Client struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// Port
type ClientRepository interface {
	FindAll() ([]Client, error)
}
