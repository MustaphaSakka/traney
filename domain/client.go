package domain

import "github.com/MustaphaSakka/traney/exception"

type Client struct {
	Id          string `db:"client_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

// Port
type ClientRepository interface {
	FindAll() ([]Client, error)
	FindById(string) (*Client, *exception.AppException)
}
