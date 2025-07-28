package services

import (
	"errors"
	"sarc/core/domain"
	interfaces "sarc/core/services/interfaces"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type resourceService struct {
	repo repositories.ResourceRepository
}

func NewResourceService(repo repositories.ResourceRepository) interfaces.ResourceService {
	return &resourceService{repo: repo}
}

func (s *resourceService) CreateResource(resource *domain.Resource) (*domain.Resource, error) {
	if err := s.repo.Create(resource); err != nil {
		return nil, err
	}
	return resource, nil
}

func (s *resourceService) GetResources() ([]domain.Resource, error) {
	return s.repo.FindAll()
}

func (s *resourceService) GetResourceByID(id uint) (*domain.Resource, error) {
	resource, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if resource == nil {
		return nil, errors.New("resource not found")
	}
	return resource, nil
}

func (s *resourceService) UpdateResource(id uint, updated *domain.Resource) (*domain.Resource, error) {
	if err := s.repo.Update(id, updated); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *resourceService) DeleteResource(id uint) error {
	return s.repo.Delete(id)
}
