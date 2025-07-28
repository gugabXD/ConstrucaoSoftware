package interfaces

import (
	"sarc/core/domain"
)

type LectureService interface {
	CreateLecture(lecture *domain.Lecture) (*domain.Lecture, error)
	GetLectures() ([]domain.Lecture, error)
	GetLectureByID(id uint) (*domain.Lecture, error)
	UpdateLecture(id uint, lecture *domain.Lecture) (*domain.Lecture, error)
	DeleteLecture(id uint) error
}
