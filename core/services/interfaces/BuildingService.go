package interfaces

import (
	"sarc/core/domain"
)

type BuildingService interface {
	CreateBuilding(building *domain.Building) (*domain.Building, error)
	GetBuildings() ([]domain.Building, error)
	GetBuildingByID(id uint) (*domain.Building, error)
	UpdateBuilding(id uint, building *domain.Building) (*domain.Building, error)
	DeleteBuilding(id uint) error
}
