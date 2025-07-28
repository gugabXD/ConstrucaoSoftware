package repositories

import "sarc/core/domain"

type DisciplineRepository interface {
	Create(discipline *domain.Discipline) error
	FindAll() ([]domain.Discipline, error)
	FindByID(id uint) (*domain.Discipline, error)
	Update(id uint, discipline *domain.Discipline) error
	Delete(id uint) error
}
