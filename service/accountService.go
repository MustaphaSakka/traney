package service

import (
	"github.com/MustaphaSakka/traney/domain"
	"github.com/MustaphaSakka/traney/dto"
	"github.com/MustaphaSakka/traney/exception"
)

const dbTSLayout = "2006-01-02 03:04:05"

type AccountService interface {
	NewAccount(dto.AccountRequest) (*dto.AccountResponse, *exception.AppException)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.AccountRequest) (*dto.AccountResponse, *exception.AppException) {
	// err := req.Validate()
	// if err != nil {
	// 	return nil, err
	// }
	// a := domain.Account{
	// 	ClientId:    req.ClientId,
	// 	OpeningDate: dbTSLayout,
	// 	AccountType: req.AccountType,
	// 	Amount:      req.Amount,
	// 	Status:      "1",
	// }
	// newAccount, err := s.repo.Save(a)
	// if err != nil {
	// 	return nil, err
	// }
	// response := newAccount.ToAccountResponseDto()

	// return &response, nil

	if err := req.Validate(); err != nil {
		return nil, err
	}
	account := domain.NewAccount(req.ClientId, req.AccountType, req.Amount)
	if newAccount, err := s.repo.Save(account); err != nil {
		return nil, err
	} else {
		return newAccount.ToAccountResponseDto(), nil
	}

}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
