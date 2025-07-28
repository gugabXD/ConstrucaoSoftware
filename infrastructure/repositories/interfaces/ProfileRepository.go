package repositories

import "sarc/core/domain"

type ProfileRepository interface {
	Create(profile *domain.Profile) error
	FindAll() ([]domain.Profile, error)
	FindByID(id uint) (*domain.Profile, error)
	Update(id uint, profile *domain.Profile) error
	Delete(id uint) error
}
