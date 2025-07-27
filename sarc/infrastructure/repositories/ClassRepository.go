package repositories

import "sarc/core/domain"

type ClassRepository interface {
	Create(class *domain.Class) error
	FindAll() ([]domain.Class, error)
	FindByID(id uint) (*domain.Class, error)
	Update(id uint, class *domain.Class) error
	Delete(id uint) error
}
