package domain

import (
	"github.com/MustaphaSakka/traney/dto"
	"github.com/MustaphaSakka/traney/exception"
)

type Account struct {
	AccountId   string
	ClientId    string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

// Port
type AccountRepository interface {
	Save(Account) (*Account, *exception.AppException)
}

func (a Account) ToAccountResponseDto() dto.AccountResponse {
	return dto.AccountResponse{AccountId: a.AccountId}
}
