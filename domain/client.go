package domain

import (
	"github.com/MustaphaSakka/traney-lib/exception"
	"github.com/MustaphaSakka/traney/dto"
)

type Client struct {
	Id          string `db:"client_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Client) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Client) ToDto() dto.ClientResponse {
	return dto.ClientResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

// Port
type ClientRepository interface {
	FindAll() ([]Client, *exception.AppException)
	FindById(string) (*Client, *exception.AppException)
}
