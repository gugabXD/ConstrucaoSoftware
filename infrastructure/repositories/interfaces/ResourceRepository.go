package repositories

import "sarc/core/domain"

type ResourceRepository interface {
	Create(resource *domain.Resource) error
	FindAll() ([]domain.Resource, error)
	FindByID(id uint) (*domain.Resource, error)
	Update(id uint, resource *domain.Resource) error
	Delete(id uint) error
}
