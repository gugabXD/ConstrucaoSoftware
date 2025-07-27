package repositories

import (
	"sarc/core/domain"

	"gorm.io/gorm"
)

type curriculumRepositoryImpl struct {
	db *gorm.DB
}

func NewCurriculumRepository(db *gorm.DB) CurriculumRepository {
	return &curriculumRepositoryImpl{db}
}

func (r *curriculumRepositoryImpl) Create(curriculum *domain.Curriculum) error {
	return r.db.Create(curriculum).Error
}

func (r *curriculumRepositoryImpl) FindAll() ([]domain.Curriculum, error) {
	var curriculums []domain.Curriculum
	err := r.db.Find(&curriculums).Error
	return curriculums, err
}

func (r *curriculumRepositoryImpl) FindByID(id uint) (*domain.Curriculum, error) {
	var curriculum domain.Curriculum
	err := r.db.First(&curriculum, id).Error
	return &curriculum, err
}

func (r *curriculumRepositoryImpl) Update(id uint, curriculum *domain.Curriculum) error {
	return r.db.Model(&domain.Curriculum{}).Where("id = ?", id).Updates(curriculum).Error
}

func (r *curriculumRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Curriculum{}, id).Error
}
