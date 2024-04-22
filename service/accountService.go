package service

import (
	"time"

	"github.com/MustaphaSakka/traney/domain"
	"github.com/MustaphaSakka/traney/dto"
	"github.com/MustaphaSakka/traney/exception"
)

type AccountService interface {
	NewAccount(dto.AccountRequest) (*dto.AccountResponse, *exception.AppException)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.AccountRequest) (*dto.AccountResponse, *exception.AppException) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		ClientId:    req.ClientId,
		OpeningDate: time.Now().Format("2006-01-02 03:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToAccountResponseDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
