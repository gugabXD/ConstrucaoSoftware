package interfaces

import (
	"sarc/core/domain"
)

type DisciplineService interface {
	CreateDiscipline(discipline *domain.Discipline) (*domain.Discipline, error)
	GetDisciplines() ([]domain.Discipline, error)
	GetDisciplineByID(id uint) (*domain.Discipline, error)
	UpdateDiscipline(id uint, discipline *domain.Discipline) (*domain.Discipline, error)
	DeleteDiscipline(id uint) error
}
