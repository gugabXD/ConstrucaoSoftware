package repositories

import (
	"sarc/core/domain"

	"gorm.io/gorm"
)

type lectureRepositoryImpl struct {
	db *gorm.DB
}

func NewLectureRepository(db *gorm.DB) LectureRepository {
	return &lectureRepositoryImpl{db}
}

func (r *lectureRepositoryImpl) Create(lecture *domain.Lecture) error {
	return r.db.Create(lecture).Error
}

func (r *lectureRepositoryImpl) FindAll() ([]domain.Lecture, error) {
	var lectures []domain.Lecture
	err := r.db.Find(&lectures).Error
	return lectures, err
}

func (r *lectureRepositoryImpl) FindByID(id uint) (*domain.Lecture, error) {
	var lecture domain.Lecture
	err := r.db.First(&lecture, id).Error
	return &lecture, err
}

func (r *lectureRepositoryImpl) Update(id uint, lecture *domain.Lecture) error {
	return r.db.Model(&domain.Lecture{}).Where("id = ?", id).Updates(lecture).Error
}

func (r *lectureRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Lecture{}, id).Error
}
