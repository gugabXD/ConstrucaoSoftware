package repositories

import (
	"sarc/core/domain"

	"gorm.io/gorm"
)

type classRepositoryImpl struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepositoryImpl{db}
}

func (r *classRepositoryImpl) Create(class *domain.Class) error {
	return r.db.Create(class).Error
}

func (r *classRepositoryImpl) FindAll() ([]domain.Class, error) {
	var classes []domain.Class
	err := r.db.Find(&classes).Error
	return classes, err
}

func (r *classRepositoryImpl) FindByID(id uint) (*domain.Class, error) {
	var class domain.Class
	err := r.db.First(&class, id).Error
	return &class, err
}

func (r *classRepositoryImpl) Update(id uint, class *domain.Class) error {
	return r.db.Model(&domain.Class{}).Where("id = ?", id).Updates(class).Error
}

func (r *classRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Class{}, id).Error
}
