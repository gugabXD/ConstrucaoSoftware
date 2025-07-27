package repositories

import (
	"sarc/core/domain"

	"gorm.io/gorm"
)

type profileRepositoryImpl struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepositoryImpl{db}
}

func (r *profileRepositoryImpl) Create(profile *domain.Profile) error {
	return r.db.Create(profile).Error
}

func (r *profileRepositoryImpl) FindAll() ([]domain.Profile, error) {
	var profiles []domain.Profile
	err := r.db.Find(&profiles).Error
	return profiles, err
}

func (r *profileRepositoryImpl) FindByID(id uint) (*domain.Profile, error) {
	var profile domain.Profile
	err := r.db.First(&profile, id).Error
	return &profile, err
}

func (r *profileRepositoryImpl) Update(id uint, profile *domain.Profile) error {
	return r.db.Model(&domain.Profile{}).Where("id = ?", id).Updates(profile).Error
}

func (r *profileRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Profile{}, id).Error
}
