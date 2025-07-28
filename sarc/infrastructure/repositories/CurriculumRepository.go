package repositories

import "sarc/core/domain"

type CurriculumRepository interface {
	Create(curriculum *domain.Curriculum) error
	FindAll() ([]domain.Curriculum, error)
	FindByID(id uint) (*domain.Curriculum, error)
	Update(id uint, curriculum *domain.Curriculum) error
	Delete(id uint) error
	AddDisciplineToCurriculum(curriculumID uint, disciplineID uint) error
}
