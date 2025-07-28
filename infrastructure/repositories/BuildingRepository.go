package repositories

import "sarc/core/domain"

type BuildingRepository interface {
	Create(building *domain.Building) error
	FindAll() ([]domain.Building, error)
	FindByID(id uint) (*domain.Building, error)
	Update(id uint, building *domain.Building) error
	Delete(id uint) error
}
