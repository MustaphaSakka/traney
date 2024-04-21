package service

import "github.com/MustaphaSakka/traney/domain"

// Port
type ClientService interface {
	GetAllClient() ([]domain.Client, error)
	GetClient(string) (*domain.Client, error)
}

type DefaultClientService struct {
	repo domain.ClientRepository
}

func (s DefaultClientService) GetAllClient() ([]domain.Client, error) {
	return s.repo.FindAll()
}

func (s DefaultClientService) GetClient(id string) (*domain.Client, error) {
	return s.repo.FindById(id)
}

func NewClientService(repository domain.ClientRepository) DefaultClientService {
	return DefaultClientService{repository}
}
