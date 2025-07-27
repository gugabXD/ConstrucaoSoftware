package repositories

import (
	"sarc/core/domain"

	"gorm.io/gorm"
)

type roomRepositoryImpl struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepositoryImpl{db}
}

func (r *roomRepositoryImpl) Create(room *domain.Room) error {
	return r.db.Create(room).Error
}

func (r *roomRepositoryImpl) FindAll() ([]domain.Room, error) {
	var rooms []domain.Room
	err := r.db.Find(&rooms).Error
	return rooms, err
}

func (r *roomRepositoryImpl) FindByID(id uint) (*domain.Room, error) {
	var room domain.Room
	err := r.db.First(&room, id).Error
	return &room, err
}

func (r *roomRepositoryImpl) Update(id uint, room *domain.Room) error {
	return r.db.Model(&domain.Room{}).Where("id = ?", id).Updates(room).Error
}

func (r *roomRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Room{}, id).Error
}
