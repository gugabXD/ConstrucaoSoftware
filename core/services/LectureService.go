package services

import (
	"errors"
	"sarc/core/domain"
	"sarc/infrastructure/repositories"
)

type LectureService interface {
	CreateLecture(lecture *domain.Lecture) (*domain.Lecture, error)
	GetLectures() ([]domain.Lecture, error)
	GetLectureByID(id uint) (*domain.Lecture, error)
	UpdateLecture(id uint, lecture *domain.Lecture) (*domain.Lecture, error)
	DeleteLecture(id uint) error
}

type lectureService struct {
	repo repositories.LectureRepository
}

func NewLectureService(repo repositories.LectureRepository) LectureService {
	return &lectureService{repo: repo}
}

func (s *lectureService) CreateLecture(lecture *domain.Lecture) (*domain.Lecture, error) {
	if err := s.repo.Create(lecture); err != nil {
		return nil, err
	}
	return lecture, nil
}

func (s *lectureService) GetLectures() ([]domain.Lecture, error) {
	return s.repo.FindAll()
}

func (s *lectureService) GetLectureByID(id uint) (*domain.Lecture, error) {
	lecture, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if lecture == nil {
		return nil, errors.New("lecture not found")
	}
	return lecture, nil
}

func (s *lectureService) UpdateLecture(id uint, updated *domain.Lecture) (*domain.Lecture, error) {
	if err := s.repo.Update(id, updated); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *lectureService) DeleteLecture(id uint) error {
	return s.repo.Delete(id)
}
