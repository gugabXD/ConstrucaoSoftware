package services

import (
	"errors"

	"sarc/core/domain"

	"gorm.io/gorm"
)

// DisciplineService interface for dependency injection and testing
type DisciplineService interface {
	CreateDiscipline(discipline *domain.Discipline) (*domain.Discipline, error)
	GetDisciplines() ([]domain.Discipline, error)
	GetDisciplineByID(id uint) (*domain.Discipline, error)
	UpdateDiscipline(id uint, discipline *domain.Discipline) (*domain.Discipline, error)
	DeleteDiscipline(id uint) error
}

type disciplineService struct {
	db *gorm.DB
}

// NewDisciplineService creates a new DisciplineService using a database connection
func NewDisciplineService(db *gorm.DB) DisciplineService {
	return &disciplineService{db: db}
}

func (s *disciplineService) CreateDiscipline(discipline *domain.Discipline) (*domain.Discipline, error) {
	if err := s.db.Create(discipline).Error; err != nil {
		return nil, err
	}
	return discipline, nil
}

func (s *disciplineService) GetDisciplines() ([]domain.Discipline, error) {
	var disciplines []domain.Discipline
	if err := s.db.Find(&disciplines).Error; err != nil {
		return nil, err
	}
	return disciplines, nil
}

func (s *disciplineService) GetDisciplineByID(id uint) (*domain.Discipline, error) {
	var discipline domain.Discipline
	if err := s.db.First(&discipline, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("discipline not found")
		}
		return nil, err
	}
	return &discipline, nil
}

func (s *disciplineService) UpdateDiscipline(id uint, updated *domain.Discipline) (*domain.Discipline, error) {
	var discipline domain.Discipline
	if err := s.db.First(&discipline, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&discipline).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &discipline, nil
}

func (s *disciplineService) DeleteDiscipline(id uint) error {
	if err := s.db.Delete(&domain.Discipline{}, id).Error; err != nil {
		return err
	}
	return nil
}
