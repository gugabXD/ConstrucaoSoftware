package services

import (
	"errors"

	"sarc/core/domain"

	"gorm.io/gorm"
)

// CurriculumService interface for dependency injection and testing
type CurriculumService interface {
	CreateCurriculum(curriculum *domain.Curriculum) (*domain.Curriculum, error)
	GetCurriculums() ([]domain.Curriculum, error)
	GetCurriculumByID(id uint) (*domain.Curriculum, error)
	UpdateCurriculum(id uint, curriculum *domain.Curriculum) (*domain.Curriculum, error)
	DeleteCurriculum(id uint) error
}

type curriculumService struct {
	db *gorm.DB
}

// NewCurriculumService creates a new CurriculumService using a database connection
func NewCurriculumService(db *gorm.DB) CurriculumService {
	return &curriculumService{db: db}
}

func (s *curriculumService) CreateCurriculum(curriculum *domain.Curriculum) (*domain.Curriculum, error) {
	if err := s.db.Create(curriculum).Error; err != nil {
		return nil, err
	}
	return curriculum, nil
}

func (s *curriculumService) GetCurriculums() ([]domain.Curriculum, error) {
	var curriculums []domain.Curriculum
	if err := s.db.Find(&curriculums).Error; err != nil {
		return nil, err
	}
	return curriculums, nil
}

func (s *curriculumService) GetCurriculumByID(id uint) (*domain.Curriculum, error) {
	var curriculum domain.Curriculum
	if err := s.db.First(&curriculum, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("curriculum not found")
		}
		return nil, err
	}
	return &curriculum, nil
}

func (s *curriculumService) UpdateCurriculum(id uint, updated *domain.Curriculum) (*domain.Curriculum, error) {
	var curriculum domain.Curriculum
	if err := s.db.First(&curriculum, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&curriculum).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &curriculum, nil
}

func (s *curriculumService) DeleteCurriculum(id uint) error {
	if err := s.db.Delete(&domain.Curriculum{}, id).Error; err != nil {
		return err
	}
	return nil
}
