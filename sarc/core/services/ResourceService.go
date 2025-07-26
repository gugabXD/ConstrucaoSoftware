package services

import (
	"errors"

	"sarc/core/domain"

	"gorm.io/gorm"
)

// ResourceService interface for dependency injection and testing
type ResourceService interface {
	CreateResource(resource *domain.Resource) (*domain.Resource, error)
	GetResources() ([]domain.Resource, error)
	GetResourceByID(id uint) (*domain.Resource, error)
	UpdateResource(id uint, resource *domain.Resource) (*domain.Resource, error)
	DeleteResource(id uint) error
}

type resourceService struct {
	db *gorm.DB
}

// NewResourceService creates a new ResourceService using a database connection
func NewResourceService(db *gorm.DB) ResourceService {
	return &resourceService{db: db}
}

func (s *resourceService) CreateResource(resource *domain.Resource) (*domain.Resource, error) {
	if err := s.db.Create(resource).Error; err != nil {
		return nil, err
	}
	return resource, nil
}

func (s *resourceService) GetResources() ([]domain.Resource, error) {
	var resources []domain.Resource
	if err := s.db.Find(&resources).Error; err != nil {
		return nil, err
	}
	return resources, nil
}

func (s *resourceService) GetResourceByID(id uint) (*domain.Resource, error) {
	var resource domain.Resource
	if err := s.db.First(&resource, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("resource not found")
		}
		return nil, err
	}
	return &resource, nil
}

func (s *resourceService) UpdateResource(id uint, updated *domain.Resource) (*domain.Resource, error) {
	var resource domain.Resource
	if err := s.db.First(&resource, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&resource).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *resourceService) DeleteResource(id uint) error {
	if err := s.db.Delete(&domain.Resource{}, id).Error; err != nil {
		return err
	}
	return nil
}
