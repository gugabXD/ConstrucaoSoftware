package services

import (
	"errors"

	"sarc/core/domain"

	"gorm.io/gorm"
)

// ClassService interface for dependency injection and testing
type ClassService interface {
	CreateClass(class *domain.Class) (*domain.Class, error)
	GetClasses() ([]domain.Class, error)
	GetClassByID(id uint) (*domain.Class, error)
	UpdateClass(id uint, class *domain.Class) (*domain.Class, error)
	DeleteClass(id uint) error
}

type classService struct {
	db *gorm.DB
}

// NewClassService creates a new ClassService using a database connection
func NewClassService(db *gorm.DB) ClassService {
	return &classService{db: db}
}

func (s *classService) CreateClass(class *domain.Class) (*domain.Class, error) {
	if err := s.db.Create(class).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (s *classService) GetClasses() ([]domain.Class, error) {
	var classes []domain.Class
	if err := s.db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (s *classService) GetClassByID(id uint) (*domain.Class, error) {
	var class domain.Class
	if err := s.db.First(&class, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("class not found")
		}
		return nil, err
	}
	return &class, nil
}

func (s *classService) UpdateClass(id uint, updated *domain.Class) (*domain.Class, error) {
	var class domain.Class
	if err := s.db.First(&class, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&class).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (s *classService) DeleteClass(id uint) error {
	if err := s.db.Delete(&domain.Class{}, id).Error; err != nil {
		return err
	}
	return nil
}
