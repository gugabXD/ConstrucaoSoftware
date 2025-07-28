package repositories

import (
	"sarc/core/domain"
)

type ResourceTypeRepository interface {
	Create(resourceType *domain.ResourceType) error
	FindAll() ([]domain.ResourceType, error)
	FindByID(id uint) (*domain.ResourceType, error)
	Update(id uint, resourceType *domain.ResourceType) error
	Delete(id uint) error
}
