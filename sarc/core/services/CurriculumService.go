package services

import (
	"errors"
	"sarc/core/domain"
	"sarc/infrastructure/repositories"
)

type CurriculumService interface {
	CreateCurriculum(curriculum *domain.Curriculum) (*domain.Curriculum, error)
	GetCurriculums() ([]domain.Curriculum, error)
	GetCurriculumByID(id uint) (*domain.Curriculum, error)
	UpdateCurriculum(id uint, curriculum *domain.Curriculum) (*domain.Curriculum, error)
	DeleteCurriculum(id uint) error
}

type curriculumService struct {
	repo repositories.CurriculumRepository
}

func NewCurriculumService(repo repositories.CurriculumRepository) CurriculumService {
	return &curriculumService{repo: repo}
}

func (s *curriculumService) CreateCurriculum(curriculum *domain.Curriculum) (*domain.Curriculum, error) {
	if err := s.repo.Create(curriculum); err != nil {
		return nil, err
	}
	return curriculum, nil
}

func (s *curriculumService) GetCurriculums() ([]domain.Curriculum, error) {
	return s.repo.FindAll()
}

func (s *curriculumService) GetCurriculumByID(id uint) (*domain.Curriculum, error) {
	curriculum, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if curriculum == nil {
		return nil, errors.New("curriculum not found")
	}
	return curriculum, nil
}

func (s *curriculumService) UpdateCurriculum(id uint, updated *domain.Curriculum) (*domain.Curriculum, error) {
	if err := s.repo.Update(id, updated); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *curriculumService) DeleteCurriculum(id uint) error {
	return s.repo.Delete(id)
}
