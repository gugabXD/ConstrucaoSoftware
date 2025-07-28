package interfaces

import (
	"sarc/core/domain"
)

type UserService interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUsers() ([]domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	UpdateUser(id uint, user *domain.User) (*domain.User, error)
	DeleteUser(id uint) error
}
