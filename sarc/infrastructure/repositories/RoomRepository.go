package repositories

import "sarc/core/domain"

type RoomRepository interface {
	Create(room *domain.Room) error
	FindAll() ([]domain.Room, error)
	FindByID(id uint) (*domain.Room, error)
	Update(id uint, room *domain.Room) error
	Delete(id uint) error
}
