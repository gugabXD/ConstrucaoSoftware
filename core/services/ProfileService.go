package services

import (
	"errors"
	"sarc/core/domain"
	"sarc/infrastructure/repositories"
)

type ProfileService interface {
	CreateProfile(profile *domain.Profile) (*domain.Profile, error)
	GetProfiles() ([]domain.Profile, error)
	GetProfileByID(id uint) (*domain.Profile, error)
	UpdateProfile(id uint, profile *domain.Profile) (*domain.Profile, error)
	DeleteProfile(id uint) error
}

type profileService struct {
	repo repositories.ProfileRepository
}

func NewProfileService(repo repositories.ProfileRepository) ProfileService {
	return &profileService{repo: repo}
}

func (s *profileService) CreateProfile(profile *domain.Profile) (*domain.Profile, error) {
	if err := s.repo.Create(profile); err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *profileService) GetProfiles() ([]domain.Profile, error) {
	return s.repo.FindAll()
}

func (s *profileService) GetProfileByID(id uint) (*domain.Profile, error) {
	profile, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, errors.New("profile not found")
	}
	return profile, nil
}

func (s *profileService) UpdateProfile(id uint, updated *domain.Profile) (*domain.Profile, error) {
	if err := s.repo.Update(id, updated); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *profileService) DeleteProfile(id uint) error {
	return s.repo.Delete(id)
}
