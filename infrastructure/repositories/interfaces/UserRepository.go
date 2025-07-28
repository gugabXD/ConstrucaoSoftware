package repositories

import "sarc/core/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindAll() ([]domain.User, error)
	FindByID(id uint) (*domain.User, error)
	Update(id uint, user *domain.User) error
	Delete(id uint) error
}
