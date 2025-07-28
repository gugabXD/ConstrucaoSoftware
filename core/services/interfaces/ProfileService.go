package interfaces

import (
	"sarc/core/domain"
)

type ProfileService interface {
	CreateProfile(profile *domain.Profile) (*domain.Profile, error)
	GetProfiles() ([]domain.Profile, error)
	GetProfileByID(id uint) (*domain.Profile, error)
	UpdateProfile(id uint, profile *domain.Profile) (*domain.Profile, error)
	DeleteProfile(id uint) error
}
