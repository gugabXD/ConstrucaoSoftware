package services

import (
	"errors"
	"sarc/core/domain"

	"gorm.io/gorm"
)

type BuildingService interface {
	CreateBuilding(building *domain.Building) (*domain.Building, error)
	GetBuildings() ([]domain.Building, error)
	GetBuildingByID(id uint) (*domain.Building, error)
	UpdateBuilding(id uint, building *domain.Building) (*domain.Building, error)
	DeleteBuilding(id uint) error
}

type buildingService struct {
	db *gorm.DB
}

func NewBuildingService(db *gorm.DB) BuildingService {
	return &buildingService{db: db}
}

func (s *buildingService) CreateBuilding(building *domain.Building) (*domain.Building, error) {
	if err := s.db.Create(building).Error; err != nil {
		return nil, err
	}
	return building, nil
}

func (s *buildingService) GetBuildings() ([]domain.Building, error) {
	var buildings []domain.Building
	if err := s.db.Find(&buildings).Error; err != nil {
		return nil, err
	}
	return buildings, nil
}

func (s *buildingService) GetBuildingByID(id uint) (*domain.Building, error) {
	var building domain.Building
	if err := s.db.First(&building, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("building not found")
		}
		return nil, err
	}
	return &building, nil
}

func (s *buildingService) UpdateBuilding(id uint, building *domain.Building) (*domain.Building, error) {
	var existing domain.Building
	if err := s.db.First(&existing, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&existing).Updates(building).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

func (s *buildingService) DeleteBuilding(id uint) error {
	if err := s.db.Delete(&domain.Building{}, id).Error; err != nil {
		return err
	}
	return nil
}
