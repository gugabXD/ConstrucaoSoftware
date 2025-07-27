package services

import (
	"errors"
	"sarc/core/domain"
	"sarc/infrastructure/repositories"
)

type ResourceService interface {
	CreateResource(resource *domain.Resource) (*domain.Resource, error)
	GetResources() ([]domain.Resource, error)
	GetResourceByID(id uint) (*domain.Resource, error)
	UpdateResource(id uint, resource *domain.Resource) (*domain.Resource, error)
	DeleteResource(id uint) error
}

type resourceService struct {
	repo repositories.ResourceRepository
}

func NewResourceService(repo repositories.ResourceRepository) ResourceService {
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
