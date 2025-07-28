package interfaces

import (
	"sarc/core/domain"
)

type ResourceService interface {
	CreateResource(resource *domain.Resource) (*domain.Resource, error)
	GetResources() ([]domain.Resource, error)
	GetResourceByID(id uint) (*domain.Resource, error)
	UpdateResource(id uint, resource *domain.Resource) (*domain.Resource, error)
	DeleteResource(id uint) error
}
