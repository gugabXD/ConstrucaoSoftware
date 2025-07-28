package services

import (
	"errors"
	"sarc/core/domain"
	"sarc/infrastructure/repositories"
)

type BuildingService interface {
	CreateBuilding(building *domain.Building) (*domain.Building, error)
	GetBuildings() ([]domain.Building, error)
	GetBuildingByID(id uint) (*domain.Building, error)
	UpdateBuilding(id uint, building *domain.Building) (*domain.Building, error)
	DeleteBuilding(id uint) error
}

type buildingService struct {
	repo repositories.BuildingRepository
}

func NewBuildingService(repo repositories.BuildingRepository) BuildingService {
	return &buildingService{repo: repo}
}

func (s *buildingService) CreateBuilding(building *domain.Building) (*domain.Building, error) {
	if err := s.repo.Create(building); err != nil {
		return nil, err
	}
	return building, nil
}

func (s *buildingService) GetBuildings() ([]domain.Building, error) {
	return s.repo.FindAll()
}

func (s *buildingService) GetBuildingByID(id uint) (*domain.Building, error) {
	building, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if building == nil {
		return nil, errors.New("building not found")
	}
	return building, nil
}

func (s *buildingService) UpdateBuilding(id uint, building *domain.Building) (*domain.Building, error) {
	if err := s.repo.Update(id, building); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *buildingService) DeleteBuilding(id uint) error {
	return s.repo.Delete(id)
}
