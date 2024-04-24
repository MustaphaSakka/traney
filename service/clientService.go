package service

import (
	"github.com/MustaphaSakka/traney-lib/exception"
	"github.com/MustaphaSakka/traney/domain"
	"github.com/MustaphaSakka/traney/dto"
)

// Port
//
//go:generate mockgen -destination=../mocks/service/mockClientService.go -package=service github.com/MustaphaSakka/traney/service ClientService
type ClientService interface {
	GetAllClient() ([]dto.ClientResponse, *exception.AppException)
	GetClient(string) (*dto.ClientResponse, *exception.AppException)
}

type DefaultClientService struct {
	repo domain.ClientRepository
}

func (s DefaultClientService) GetAllClient() ([]dto.ClientResponse, *exception.AppException) {
	clients, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.ClientResponse, 0)
	for _, c := range clients {
		response = append(response, c.ToDto())
	}

	return response, err
}

func (s DefaultClientService) GetClient(id string) (*dto.ClientResponse, *exception.AppException) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil

}

func NewClientService(repository domain.ClientRepository) DefaultClientService {
	return DefaultClientService{repository}
}
