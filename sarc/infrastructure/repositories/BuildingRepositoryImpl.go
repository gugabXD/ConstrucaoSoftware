package repositories

import (
	"sarc/core/domain"

	"gorm.io/gorm"
)

type buildingRepositoryImpl struct {
	db *gorm.DB
}

func NewBuildingRepository(db *gorm.DB) BuildingRepository {
	return &buildingRepositoryImpl{db}
}

func (r *buildingRepositoryImpl) Create(building *domain.Building) error {
	return r.db.Create(building).Error
}

func (r *buildingRepositoryImpl) FindAll() ([]domain.Building, error) {
	var buildings []domain.Building
	err := r.db.Find(&buildings).Error
	return buildings, err
}

func (r *buildingRepositoryImpl) FindByID(id uint) (*domain.Building, error) {
	var building domain.Building
	err := r.db.First(&building, id).Error
	return &building, err
}

func (r *buildingRepositoryImpl) Update(id uint, building *domain.Building) error {
	return r.db.Model(&domain.Building{}).Where("id = ?", id).Updates(building).Error
}

func (r *buildingRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Building{}, id).Error
}
