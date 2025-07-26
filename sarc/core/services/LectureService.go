package services

import (
	"errors"

	"sarc/core/domain"

	"gorm.io/gorm"
)

// LectureService interface for dependency injection and testing
type LectureService interface {
	CreateLecture(lecture *domain.Lecture) (*domain.Lecture, error)
	GetLectures() ([]domain.Lecture, error)
	GetLectureByID(id uint) (*domain.Lecture, error)
	UpdateLecture(id uint, lecture *domain.Lecture) (*domain.Lecture, error)
	DeleteLecture(id uint) error
}

type lectureService struct {
	db *gorm.DB
}

// NewLectureService creates a new LectureService using a database connection
func NewLectureService(db *gorm.DB) LectureService {
	return &lectureService{db: db}
}

func (s *lectureService) CreateLecture(lecture *domain.Lecture) (*domain.Lecture, error) {
	if err := s.db.Create(lecture).Error; err != nil {
		return nil, err
	}
	return lecture, nil
}

func (s *lectureService) GetLectures() ([]domain.Lecture, error) {
	var lectures []domain.Lecture
	if err := s.db.Find(&lectures).Error; err != nil {
		return nil, err
	}
	return lectures, nil
}

func (s *lectureService) GetLectureByID(id uint) (*domain.Lecture, error) {
	var lecture domain.Lecture
	if err := s.db.First(&lecture, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lecture not found")
		}
		return nil, err
	}
	return &lecture, nil
}

func (s *lectureService) UpdateLecture(id uint, updated *domain.Lecture) (*domain.Lecture, error) {
	var lecture domain.Lecture
	if err := s.db.First(&lecture, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&lecture).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &lecture, nil
}

func (s *lectureService) DeleteLecture(id uint) error {
	if err := s.db.Delete(&domain.Lecture{}, id).Error; err != nil {
		return err
	}
	return nil
}
