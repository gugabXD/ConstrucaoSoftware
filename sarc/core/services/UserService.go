package services

import (
	"errors"

	"sarc/core/domain"

	"gorm.io/gorm"
)

// UserService interface for dependency injection and testing
type UserService interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUsers() ([]domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	UpdateUser(id uint, user *domain.User) (*domain.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	db *gorm.DB
}

// NewUserService creates a new UserService using a database connection
func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) CreateUser(user *domain.User) (*domain.User, error) {
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUsers() ([]domain.User, error) {
	var users []domain.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(id uint, updated *domain.User) (*domain.User, error) {
	var user domain.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&user).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) DeleteUser(id uint) error {
	if err := s.db.Delete(&domain.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
