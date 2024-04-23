package domain

import (
	"github.com/MustaphaSakka/traney/dto"
	"github.com/MustaphaSakka/traney/exception"
)

const dbTSLayout = "2006-01-02 03:04:05"

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

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/MustaphaSakka/traney/domain AccountRepository
func (a Account) ToAccountResponseDto() *dto.AccountResponse {
	return &dto.AccountResponse{AccountId: a.AccountId}
}

func NewAccount(clientId, accountType string, amount float64) Account {
	return Account{
		ClientId:    clientId,
		OpeningDate: dbTSLayout,
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}
