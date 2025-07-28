package services

import (
	"errors"
	"sarc/core/domain"
	interfaces "sarc/core/services/interfaces"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type classService struct {
	repo repositories.ClassRepository
}

func NewClassService(repo repositories.ClassRepository) interfaces.ClassService {
	return &classService{repo: repo}
}

func (s *classService) CreateClass(class *domain.Class) (*domain.Class, error) {
	if err := s.repo.Create(class); err != nil {
		return nil, err
	}
	return class, nil
}

func (s *classService) GetClasses() ([]domain.Class, error) {
	return s.repo.FindAll()
}

func (s *classService) GetClassByID(id uint) (*domain.Class, error) {
	class, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if class == nil {
		return nil, errors.New("class not found")
	}
	return class, nil
}

func (s *classService) UpdateClass(id uint, class *domain.Class) (*domain.Class, error) {
	if err := s.repo.Update(id, class); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *classService) DeleteClass(id uint) error {
	return s.repo.Delete(id)
}
