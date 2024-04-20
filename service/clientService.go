package service

import "github.com/MustaphaSakka/traney/domain"

// Port
type ClientService interface {
	GetAllClient() ([]domain.Client, error)
}

type DefaultClientService struct {
	repo domain.ClientRepository
}

func (s DefaultClientService) GetAllClient() ([]domain.Client, error) {
	return s.repo.FindAll()
}

func NewClientService(repository domain.ClientRepository) DefaultClientService {
	return DefaultClientService{repository}
}
