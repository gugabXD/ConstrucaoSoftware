package repositories

import (
	"sarc/core/domain"

	"gorm.io/gorm"
)

type resourceRepositoryImpl struct {
	db *gorm.DB
}

func NewResourceRepository(db *gorm.DB) ResourceRepository {
	return &resourceRepositoryImpl{db}
}

func (r *resourceRepositoryImpl) Create(resource *domain.Resource) error {
	return r.db.Create(resource).Error
}

func (r *resourceRepositoryImpl) FindAll() ([]domain.Resource, error) {
	var resources []domain.Resource
	err := r.db.Find(&resources).Error
	return resources, err
}

func (r *resourceRepositoryImpl) FindByID(id uint) (*domain.Resource, error) {
	var resource domain.Resource
	err := r.db.First(&resource, id).Error
	return &resource, err
}

func (r *resourceRepositoryImpl) Update(id uint, resource *domain.Resource) error {
	return r.db.Model(&domain.Resource{}).Where("id = ?", id).Updates(resource).Error
}

func (r *resourceRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Resource{}, id).Error
}
