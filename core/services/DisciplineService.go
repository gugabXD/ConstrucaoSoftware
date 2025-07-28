package services

import (
	"errors"
	"sarc/core/domain"
	interfaces "sarc/core/services/interfaces"
	repositories "sarc/infrastructure/repositories/interfaces"
)

type disciplineService struct {
	repo repositories.DisciplineRepository
}

func NewDisciplineService(repo repositories.DisciplineRepository) interfaces.DisciplineService {
	return &disciplineService{repo: repo}
}

func (s *disciplineService) CreateDiscipline(discipline *domain.Discipline) (*domain.Discipline, error) {
	if err := s.repo.Create(discipline); err != nil {
		return nil, err
	}
	return discipline, nil
}

func (s *disciplineService) GetDisciplines() ([]domain.Discipline, error) {
	return s.repo.FindAll()
}

func (s *disciplineService) GetDisciplineByID(id uint) (*domain.Discipline, error) {
	discipline, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if discipline == nil {
		return nil, errors.New("discipline not found")
	}
	return discipline, nil
}

func (s *disciplineService) UpdateDiscipline(id uint, updated *domain.Discipline) (*domain.Discipline, error) {
	if err := s.repo.Update(id, updated); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *disciplineService) DeleteDiscipline(id uint) error {
	return s.repo.Delete(id)
}
