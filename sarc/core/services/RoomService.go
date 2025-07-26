package services

import (
	"errors"

	"sarc/core/domain"

	"gorm.io/gorm"
)

// RoomService interface for dependency injection and testing
type RoomService interface {
	CreateRoom(room *domain.Room) (*domain.Room, error)
	GetRooms() ([]domain.Room, error)
	GetRoomByID(id uint) (*domain.Room, error)
	UpdateRoom(id uint, room *domain.Room) (*domain.Room, error)
	DeleteRoom(id uint) error
}

type roomService struct {
	db *gorm.DB
}

// NewRoomService creates a new RoomService using a database connection
func NewRoomService(db *gorm.DB) RoomService {
	return &roomService{db: db}
}

func (s *roomService) CreateRoom(room *domain.Room) (*domain.Room, error) {
	if err := s.db.Create(room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func (s *roomService) GetRooms() ([]domain.Room, error) {
	var rooms []domain.Room
	if err := s.db.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *roomService) GetRoomByID(id uint) (*domain.Room, error) {
	var room domain.Room
	if err := s.db.First(&room, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}
	return &room, nil
}

func (s *roomService) UpdateRoom(id uint, updated *domain.Room) (*domain.Room, error) {
	var room domain.Room
	if err := s.db.First(&room, id).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&room).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (s *roomService) DeleteRoom(id uint) error {
	if err := s.db.Delete(&domain.Room{}, id).Error; err != nil {
		return err
	}
	return nil
}
