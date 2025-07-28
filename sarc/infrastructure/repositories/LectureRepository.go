package repositories

import "sarc/core/domain"

type LectureRepository interface {
	Create(lecture *domain.Lecture) error
	FindAll() ([]domain.Lecture, error)
	FindByID(id uint) (*domain.Lecture, error)
	Update(id uint, lecture *domain.Lecture) error
	Delete(id uint) error
}
