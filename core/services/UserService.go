package services

import (
	"errors"
	"sarc/core/domain"
	interfaces "sarc/core/services/interfaces"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) interfaces.UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *domain.User) (*domain.User, error) {
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) UpdateUser(id uint, updated *domain.User) (*domain.User, error) {
	if err := s.repo.Update(id, updated); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
