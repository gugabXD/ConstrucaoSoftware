package interfaces

import (
	"sarc/core/domain"
)

type RoomService interface {
	CreateRoom(room *domain.Room) (*domain.Room, error)
	GetRooms() ([]domain.Room, error)
	GetRoomByID(id uint) (*domain.Room, error)
	UpdateRoom(id uint, room *domain.Room) (*domain.Room, error)
	DeleteRoom(id uint) error
}
