package interfaces

import "sarc/core/domain"

type ClassService interface {
	CreateClass(class *domain.Class) (*domain.Class, error)
	GetClasses() ([]domain.Class, error)
	GetClassByID(id uint) (*domain.Class, error)
	UpdateClass(id uint, class *domain.Class) (*domain.Class, error)
	DeleteClass(id uint) error
}
