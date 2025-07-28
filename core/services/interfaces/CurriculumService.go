package interfaces

import (
	"sarc/core/domain"
)

type CurriculumService interface {
	CreateCurriculum(curriculum *domain.Curriculum) (*domain.Curriculum, error)
	GetCurriculums() ([]domain.Curriculum, error)
	GetCurriculumByID(id uint) (*domain.Curriculum, error)
	UpdateCurriculum(id uint, curriculum *domain.Curriculum) (*domain.Curriculum, error)
	DeleteCurriculum(id uint) error
	AddDisciplineToCurriculum(curriculumID uint, disciplineID uint) error
}
