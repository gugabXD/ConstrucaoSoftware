package services

import (
	"errors"

	"sarc/core/domain"

	"gorm.io/gorm"
)

// ProfileService interface for dependency injection and testing
type ProfileService interface {
	CreateProfile(profile *domain.Profile) (*domain.Profile, error)
	GetProfiles() ([]domain.Profile, error)
	GetProfileByID(id uint) (*domain.Profile, error)
	UpdateProfile(id uint, profile *domain.Profile) (*domain.Profile, error)
	DeleteProfile(id uint) error
}

type profileService struct {
	db *gorm.DB
}

// NewProfileService creates a new ProfileService using a database connection
func NewProfileService(db *gorm.DB) ProfileService {
	return &profileService{db: db}
}

func (s *profileService) CreateProfile(profile *domain.Profile) (*domain.Profile, error) {
	if err := s.db.Create(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *profileService) GetProfiles() ([]domain.Profile, error) {
	var profiles []domain.Profile
	if err := s.db.Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}

func (s *profileService) GetProfileByID(id uint) (*domain.Profile, error) {
	var profile domain.Profile
	if err := s.db.First(&profile, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}
	return &profile, nil
}

func (s *profileService) UpdateProfile(id uint, updated *domain.Profile) (*domain.Profile, error) {
	var profile domain.Profile
	if err := s.db.First(&profile, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&profile).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *profileService) DeleteProfile(id uint) error {
	if err := s.db.Delete(&domain.Profile{}, id).Error; err != nil {
		return err
	}
	return nil
}
