package repositories

import (
	"sarc/core/domain"

	"gorm.io/gorm"
)

type disciplineRepositoryImpl struct {
	db *gorm.DB
}

func NewDisciplineRepository(db *gorm.DB) DisciplineRepository {
	return &disciplineRepositoryImpl{db}
}

func (r *disciplineRepositoryImpl) Create(discipline *domain.Discipline) error {
	return r.db.Create(discipline).Error
}

func (r *disciplineRepositoryImpl) FindAll() ([]domain.Discipline, error) {
	var disciplines []domain.Discipline
	err := r.db.Find(&disciplines).Error
	return disciplines, err
}

func (r *disciplineRepositoryImpl) FindByID(id uint) (*domain.Discipline, error) {
	var discipline domain.Discipline
	err := r.db.First(&discipline, id).Error
	return &discipline, err
}

func (r *disciplineRepositoryImpl) Update(id uint, discipline *domain.Discipline) error {
	return r.db.Model(&domain.Discipline{}).Where("id = ?", id).Updates(discipline).Error
}

func (r *disciplineRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Discipline{}, id).Error
}
